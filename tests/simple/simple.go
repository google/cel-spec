package simple

import (
	"fmt"

	"github.com/golang/protobuf/proto"

	ctpb "github.com/google/cel-spec/proto/test/v1/conformanceTest"
	exprbp "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

var (
	trueval = &exprpb.Value{ Kind: &exprpb.Value_BoolValue{true} }
)

// Match checks the expectation in the response matcher against the
// actual result of evaluation.  Returns nil if the expectation
// matches the actual, otherwise returns an error describing the difference.
// Calling this function implies that the interpretation succeeded
// in the parse and check phases.  See MatchValue() for the normalization
// applied to values for matching.
func Match(m *ctpb.EvalResponseMatcher, actual *exprpb.ExprValue) error {
	switch m.Kind.(type) {
	case *ctpb.EvalResponseMatcher_Value:
		switch actual.Kind.(type) {
		case *exprpb.ExprValue_Value:
			return MatchValue(m.GetValue(), actual.GetValue())
		}
		return fmt.Errorf("Expected value, got %v", actual)
	case *ctpb.EvalResponseMatcher_Errors:
		switch actual.Kind.(type) {
		case *exprpb.ExprValue_Error:
			// TODO match errors
			return nil
		}
		return fmt.Errorf("Expected error, got %v", actual)
	case *ctpb.EvalResponseMatcher_Unknowns:
		switch actual.Kind.(type) {
		case *exprpb.ExprValue_Error:
			// TODO match unknowns
			return nil
		}
		return fmt.Errorf("Expected unknown, got %v", actual)
	case *ctpb.EvalResponseMatcher_ParseFailureRegexp:
		return fmt.Errorf("parse succeeded but expected failure: %v", actual)
	case *ctpb.EvalResponseMatcher_CheckFailureRegexp:
		return fmt.Errorf("check succeeded but expected failure: %v", actual)
	case *ctpb.EvalResponseMatcher_Trueval:
		switch actual.Kind.(type) {
		case *exprpb.ExprValue_Value:
			return MatchValue(trueval, actual.GetValue())
		}
		return fmt.Errorf("Expected true, got %v", actual)
	}
	return err("Unsupported matcher kind")
}

// MatchValue returns whether the actual value is equal to the
// expected value, modulo the following normalization:
// 1) All floating-point NaN values are equal.
// 2) Map comparisons ignore order.
func MatchValue(expected *exprpb.Value, actual *exprpb.Value) error {
	// XXX for now, just compare the protos.
	if !proto.Equal(expected, actual) {
		return fmt.Errorf("Expected [%v], Actual [%v]", expected, actual
	}
	return nil
}

type runConfig struct {
	parseClient *celrpc.Celclient,
	checkClient *celrpc.Celclient,
	evalClient *celrpc.Celclient,
}

func (r *runConfig) RunTest(t *ctpb.SimpleEvalTest) error {
	err := ValidateTest(t)
	if err != nil {
		return err
	}
	m := t.Expected

	// Parse
	preq := cs.ParseRequest{
		CelSource: ctpb.Expr,
		SourceLocation: ctpb.Name,
		EnableMacros: ctpb.EnableMacros,
	}
	pres, err := r.parseClient.Parse(context.Background(), &preq)
	if err != nil {
		return fmt.Errorf("%s: Fatal parse RPC error: %v", t.Name, err)
	}
	if pres == nil {
		return fmt.Errorf("%s: Empty parse RPC response", t.Name)
	}
	parsedExpr := pres.Expr
	if parsedExpr == nil {
		switch t.Expected.Kind.(type) {
		case *ctpb.Matcher_ParseFailureRegexp:
			// TODO interpret regexp
			return nil
		}
		return fmt.Errorf("%s: Fatal parse errors: %v", t.Name, pres.Fatals)
	}
	switch t.Expected.Kind.(type) {
	case *ctpb.Matcher_ParseFailureRegexp:
		return fmt.Errorf("%s: parse succeeded but expected failure: %v", t.Name)
	}
	if parsedExpr.Expr == nil {
		return fmt.Errorf("%s: Empty root expression", t.Name)
	}
	rootId := parsedExpr.Expr.Id

	// Check (optional)
	var checkedExpr *checked.CheckedExpr
	if t.EnableCheck {
		creq := cs.CheckRequest{
			Expr: parsedExpr,
			TypeEnv: t.TypeEnv,
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
			switch t.Expected.Kind.(type) {
			case *ctpb.Matcher_CheckFailureRegexp:
				// TODO interpret regexp
				return nil
			}
			return fmt.Errorf("%s: Fatal check errors: %v", t.Name, cres.Fatals)
		}
		switch t.Expected.Kind.(type) {
		case *ctpb.Matcher_CheckFailureRegexp:
			return fmt.Errorf("%s: check succeeded but expected failure: %v", t.Name, checkedExpr)
		}
		topType, present := checkedExpr.TypeMap[rootId]
		if !present {
			return fmt.Errorf("%s: No type for top level expression: %v", t.Name, cres)
		}
	}

	// Eval
	var ereq cs.EvalRequest
	if checkedExpr == nil {
		ereq = cs.EvalRequest{
			ExprKind: &cs.EvalRequest_Parsed{parsedExpr},
			Bindings: t.Bindings,
		}
	} else {
		ereq = cs.EvalRequest{
			ExprKind: &cs.EvalRequest_Checked{checkedExpr},
			Bindings: t.Bindings,
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

func ValidateTest(t *ctpb.SimpleEvalTest) error {
	if t.Name == nil {
		return fmt.Errorf("Simple test has no name")
	}
	if t.Expr == nil {
		return fmt.Errorf("%s: no expression", t.Name)
	}
	if t.Expected == nil {
		return fmt.Errorf("%s: no expected result", t.Name)
	}
	if t.Expected.GetKind() == nil {
		return fmt.Errorf("%s: no expected result kind", t.Name)
	}
	return nil
}
