package tests

import (
	"github.com/golang/protobuf/proto"

	ctpb "github.com/google/cel-spec/proto/test/v1/conformanceTest"
	exprbp "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

func Match(m *ctpb.EvalResponseMatcher, actual *exprpb.ExprValue) bool {
	switch m.Kind.(type) {
	case *ctpb.EvalResponseMatcher_Value:
		expected := m.GetValue()
		switch actual.Kind.(type) {
		case *exprpb.ExprValue_Value:
			return MatchVlaue(expected, actual.GetValue())
		case *exprpb.ExprValue_Error:
		case *exprpb.ExprValue_Unknown:
		}
	case *ctpb.EvalResponseMatcher_Errors:
	case *ctpb.EvalResponseMatcher_Unknowns:
	case *ctpb.EvalResponseMatcher_ParseFailureRegexp:
	case *ctpb.EvalResponseMatcher_CheckFailureRegexp:
	case *ctpb.EvalResponseMatcher_Trueval:
	default:
		return err("Unsupported matcher kind")
	}
}

// MatchValue returns whether the actual value is equal to the
// expected value, modulo the following normalization:
// 1) All floating-point NaN values are equal.
// 2) Map comparisons ignore order.
func MatchValue(expected *exprpb.Value, actual *exprpb.Value) bool {
	// XXX for now, just compare the protos.
	return proto.Equal(expected, actual)

}

func RunZZ
