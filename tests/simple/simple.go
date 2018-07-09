package simple

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/google/cel-spec/tools/celrpc"

	ctpb "github.com/google/cel-spec/proto/test/v1/conformanceTest"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
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
	case *ctpb.EvalResponseMatcher_ParseFailureRegex:
		return fmt.Errorf("parse succeeded but expected failure: %v", actual)
	case *ctpb.EvalResponseMatcher_CheckFailureRegex:
		return fmt.Errorf("check succeeded but expected failure: %v", actual)
	case *ctpb.EvalResponseMatcher_Trueval:
		switch actual.Kind.(type) {
		case *exprpb.ExprValue_Value:
			return MatchValue(trueval, actual.GetValue())
		}
		return fmt.Errorf("Expected true, got %v", actual)
	}
	return fmt.Errorf("Unsupported matcher kind")
}

// MatchValue returns whether the actual value is equal to the
// expected value, modulo the following normalization:
// 1) All floating-point NaN values are equal.
// 2) Map comparisons ignore order.
func MatchValue(expected *exprpb.Value, actual *exprpb.Value) error {
	// XXX for now, just compare the protos.
	if !proto.Equal(expected, actual) {
		return fmt.Errorf("Expected [%v], Actual [%v]", expected, actual)
	}
	return nil
}

type runConfig struct {
	parseClient *celrpc.ConfClient
	checkClient *celrpc.ConfClient
	evalClient *celrpc.ConfClient
}

func (r *runConfig) RunTest(t *ctpb.SimpleEvalTest) error {
	err := ValidateTest(t)
	if err != nil {
		return err
	}
	m := t.Expected

	// Parse
	preq := exprpb.ParseRequest{
		CelSource: t.Expr,
		SourceLocation: t.Name,
		DisableMacros: t.DisableMacros,
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
		switch m.Kind.(type) {
		case *ctpb.EvalResponseMatcher_ParseFailureRegex:
			// TODO interpret regex
			return nil
		}
		return fmt.Errorf("%s: Fatal parse errors: %v", t.Name, pres.Issues)
	}
	switch m.Kind.(type) {
	case *ctpb.EvalResponseMatcher_ParseFailureRegex:
		return fmt.Errorf("%s: parse succeeded but expected failure: %v", t.Name)
	}
	if parsedExpr.Expr == nil {
		return fmt.Errorf("%s: Empty root expression", t.Name)
	}
	rootId := parsedExpr.Expr.Id

	// Check (optional)
	var checkedExpr *exprpb.CheckedExpr
	if t.EnableCheck {
		creq := exprpb.CheckRequest{
			ParsedExpr: parsedExpr,
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
			switch m.Kind.(type) {
			case *ctpb.EvalResponseMatcher_CheckFailureRegex:
				// TODO interpret regex
				return nil
			}
			return fmt.Errorf("%s: Fatal check errors: %v", t.Name, cres.Issues)
		}
		switch m.Kind.(type) {
		case *ctpb.EvalResponseMatcher_CheckFailureRegex:
			return fmt.Errorf("%s: check succeeded but expected failure: %v", t.Name, checkedExpr)
		}
		_, present := checkedExpr.TypeMap[rootId]
		if !present {
			return fmt.Errorf("%s: No type for top level expression: %v", t.Name, cres)
		}
	}

	// Eval
	var ereq exprpb.EvalRequest
	if checkedExpr == nil {
		ereq = exprpb.EvalRequest{
			ExprKind: &exprpb.EvalRequest_ParsedExpr{parsedExpr},
			Bindings: t.Bindings,
		}
	} else {
		ereq = exprpb.EvalRequest{
			ExprKind: &exprpb.EvalRequest_CheckedExpr{checkedExpr},
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
	return Match(m, eres.Result)
}

func ValidateTest(t *ctpb.SimpleEvalTest) error {
	if t.Name == "" {
		return fmt.Errorf("Simple test has no name")
	}
	if t.Expr == "" {
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
