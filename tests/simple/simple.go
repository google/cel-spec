/*
Package simple runs end-to-end CEL conformance tests against
ConformanceService servers.  The "simple" tests run the Parse /
Check (optional) / Eval pipeline and compare the result against an
expected value, error, or unknown from the Eval phase.  To validate the
intermediate results from the Parse or Check phases, use a different
test driver.

Each phase can be sent to a different ConformanceService server.  Thus a
partial implementation can be tested by using other implementations for
the missing phases.  This also validates the interoperativity.

Example test data:

	name: "basic"
	description: "Basic tests that all implementations should pass."
	section {
	  name: "self_eval"
	  description: "Simple self-evaluating forms."
	  test {
	    name: "self_eval_zero"
	    expr: "0"
	    value: { int64_value: 0 }
	  }
	}
	section {
	  name: "arithmetic"
	  description: "Numeric arithmetic checks."
	  test {
	    name: "one plus one"
	    description: "Uses implicit match against 'true'."
	    expr: "1 + 1 == 2"
	  }
	}

*/
package simple

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/google/cel-spec/tools/celrpc"

	spb "github.com/google/cel-spec/proto/test/v1/simple"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

var (
	trueval = &exprpb.Value{Kind: &exprpb.Value_BoolValue{BoolValue: true}}
)

// Match checks the expectation in the result matcher against the
// actual result of evaluation.  Returns nil if the expectation
// matches the actual, otherwise returns an error describing the difference.
// Calling this function implies that the interpretation succeeded
// in the parse and check phases.  See MatchValue() for the normalization
// applied to values for matching.
func Match(t *spb.SimpleTest, actual *exprpb.ExprValue) error {
	switch t.ResultMatcher.(type) {
	case *spb.SimpleTest_Value:
		want := t.GetValue()
		switch actual.Kind.(type) {
		case *exprpb.ExprValue_Value:
			return MatchValue(t.Name, want, actual.GetValue())
		}
		return fmt.Errorf("Got %v, want value %v", actual, want)
	case *spb.SimpleTest_EvalError:
		switch actual.Kind.(type) {
		case *exprpb.ExprValue_Error:
			// TODO match errors
			return nil
		}
		return fmt.Errorf("Got %v, want error", actual)
	// TODO support any_eval_errors
	case *spb.SimpleTest_Unknown:
		switch actual.Kind.(type) {
		case *exprpb.ExprValue_Error:
			// TODO match unknowns
			return nil
		}
		return fmt.Errorf("Got %v, want unknown", actual)
	// TODO support any_unknowns
	case nil:
		// Defaults to a match against a true value.
		switch actual.Kind.(type) {
		case *exprpb.ExprValue_Value:
			return MatchValue(t.Name, trueval, actual.GetValue())
		}
		return fmt.Errorf("Got %v, want true", actual)
	}
	return fmt.Errorf("Unsupported matcher kind")
}

// MatchValue returns whether the actual value is equal to the
// expected value, modulo the following normalization:
//	1) All floating-point NaN values are equal.
//	2) Map comparisons ignore order.
func MatchValue(tag string, expected *exprpb.Value, actual *exprpb.Value) error {
	// TODO write normalized comparator.
	switch expected.GetKind().(type) {
	case *exprpb.Value_MapValue:
		// Maps are handled as repeated entries, but the entries need to be
		// compared using set equality semantics.
		expectedMap := expected.GetMapValue()
		actualMap := actual.GetMapValue()
		if actualMap == nil || expectedMap == nil {
			break
		}
		expectedEntries := expectedMap.GetEntries()
		actualEntries := actualMap.GetEntries()
		if len(expectedEntries) != len(actualEntries) {
			break
		}
		for _, expectedElem := range expectedEntries {
			found := false
			for _, actualElem := range actualEntries {
				keyErr := MatchValue(tag, expectedElem.GetKey(), actualElem.GetKey())
				if keyErr != nil {
					continue
				}
				valErr := MatchValue(tag, expectedElem.GetValue(), actualElem.GetValue())
				if valErr == nil {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("%s: Eval got [%v], want [%v]", tag, actual, expected)
			}
		}
	default:
		// By default, just compare the protos.
		if !proto.Equal(expected, actual) {
			return fmt.Errorf("%s: Eval got [%v], want [%v]", tag, actual, expected)
		}
	}
	return nil
}

// runConfig holds client stubs for the servers to use
// for the various phases.  Some phases might use the
// same server.
type runConfig struct {
	parseClient *celrpc.ConfClient
	checkClient *celrpc.ConfClient
	evalClient  *celrpc.ConfClient
}

// RunTest runs the test described by t, returning an error for any
// violation of expectations.
func (r *runConfig) RunTest(t *spb.SimpleTest) error {
	err := ValidateTest(t)
	if err != nil {
		return err
	}

	// Parse
	preq := exprpb.ParseRequest{
		CelSource:      t.Expr,
		SourceLocation: t.Name,
		DisableMacros:  t.DisableMacros,
	}
	pres, err := r.parseClient.Parse(context.Background(), &preq)
	if err != nil {
		return fmt.Errorf("%s: Fatal parse RPC error: %v", t.Name, err)
	}
	if pres == nil {
		return fmt.Errorf("%s: Empty parse RPC response", t.Name)
	}
	parsedExpr := pres.ParsedExpr
	if parsedExpr == nil {
		return fmt.Errorf("%s: Fatal parse errors: %v", t.Name, pres.Issues)
	}
	if parsedExpr.Expr == nil {
		return fmt.Errorf("%s: parse returned empty root expression", t.Name)
	}
	rootID := parsedExpr.Expr.Id

	// Check (optional)
	var checkedExpr *exprpb.CheckedExpr
	if !t.DisableCheck {
		creq := exprpb.CheckRequest{
			ParsedExpr: parsedExpr,
			TypeEnv:    t.TypeEnv,
			Container:  t.Container,
		}
		cres, err := r.checkClient.Check(context.Background(), &creq)
		if err != nil {
			return fmt.Errorf("%s: Fatal check RPC error: %v", t.Name, err)
		}
		if cres == nil {
			return fmt.Errorf("%s: Empty check RPC response", t.Name)
		}
		checkedExpr = cres.CheckedExpr
		if checkedExpr == nil {
			return fmt.Errorf("%s: Fatal check errors: %v", t.Name, cres.Issues)
		}
		_, present := checkedExpr.TypeMap[rootID]
		if !present {
			return fmt.Errorf("%s: No type for top level expression: %v", t.Name, cres)
		}
		// TODO: validate that the inferred type is compatible
		// with the expected value, if any, in the eval matcher.
	}
	// TODO: if check phase disabled, check anyhow and ensure it fails.

	// Eval
	var ereq exprpb.EvalRequest
	if checkedExpr == nil {
		ereq = exprpb.EvalRequest{
			ExprKind:  &exprpb.EvalRequest_ParsedExpr{ParsedExpr: parsedExpr},
			Bindings:  t.Bindings,
			Container: t.Container,
		}
	} else {
		ereq = exprpb.EvalRequest{
			ExprKind:  &exprpb.EvalRequest_CheckedExpr{CheckedExpr: checkedExpr},
			Bindings:  t.Bindings,
			Container: t.Container,
		}
	}
	eres, err := r.evalClient.Eval(context.Background(), &ereq)
	if err != nil {
		return fmt.Errorf("%s: Fatal eval RPC error: %v", t.Name, err)
	}
	if eres == nil || eres.Result == nil {
		return fmt.Errorf("%s: empty eval response", t.Name)
	}
	return Match(t, eres.Result)
}

// ValidateTest checks whether a simple test has the required fields.
func ValidateTest(t *spb.SimpleTest) error {
	if t.Name == "" {
		return fmt.Errorf("Simple test has no name")
	}
	if t.Expr == "" {
		return fmt.Errorf("%s: no expression", t.Name)
	}
	return nil
}
