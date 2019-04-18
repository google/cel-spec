/*
Package envcheck runs CEL conformance tests to check that the
runtime supports a set of functions.  A set of checker declarations
is scanned to produce a set of CEL expressions, each of which is then
compiled and sent to the runtime.

Identifier declarations are compiled to an expression of just that
identifiers.  For instance, the "int" type identifier produces:

	int

Function declarations are compiled to a separate expression for each
overload.  The expression is an invocation of the overload with "zeroish"
arguments of the appropriate type.  The zeroish arguments are:

	int		0
	uint		0u
	double		0.0
	bool		false
	string		""
	bytes		b""
	null_type	null
	type		type
	list<A>		[]
	map<A,B>	{}
	enum E		0
	message M	M{}

For instance, the "_/_" function with overloads

	_/_: (int, int) -> int
	_/_: (uint, uint) -> uint
	_/_: (double, double) -> double

compiles to the expressions

	(0)/(0)
	(0u)/(0u)
	(0.0)/(0.0)

which are then evaluated.

This test suite does not check that the overloads are implemented correctly,
only that they are implemented at all.  The test will pass unless the
expression evaluates (with no bindings) to any result or error other than
"no_matching_overload".  For instance, the first two expressions for _/_
will generate division-by-zero errors, but this will pass the test.
*/
package envcheck

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/cel-spec/tools/celrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	spb "google.golang.org/genproto/googleapis/rpc/status"
)

// runConfig holds client stubs for the servers to use
// for the various phases.  Some phases might use the
// same server.
type runConfig struct {
	parseClient *celrpc.ConfClient
	evalClient  *celrpc.ConfClient
}

var (
	noOverload = &exprpb.ExprValue{Kind: &exprpb.ExprValue_Error{
		Error: &exprpb.ErrorSet{Errors: []*spb.Status{
			status.New(codes.InvalidArgument, "no_matching_overload").Proto(),
		}},
	}}
)

// A verifier returns true if the overload or identifier is supported.
type verifier func(*exprpb.ExprValue) bool

func verifyIdentifier(result *exprpb.ExprValue) bool {
	_, ok := result.Kind.(*exprpb.ExprValue_Value)
	return ok
}

func verifyOverload(result *exprpb.ExprValue) bool {
	return !proto.Equal(result, noOverload)
}

func zeroValuePrimitive(p exprpb.Type_PrimitiveType) (string, error) {
	switch p {
	case exprpb.Type_BOOL:
		return "false", nil
	case exprpb.Type_INT64:
		return "0", nil
	case exprpb.Type_UINT64:
		return "0u", nil
	case exprpb.Type_DOUBLE:
		return "0.0", nil
	case exprpb.Type_STRING:
		return "''", nil
	case exprpb.Type_BYTES:
		return "b''", nil
	default:
		return "", fmt.Errorf("unknown primitive type: %v", p)
	}
}

func zeroValueWellKnown(w exprpb.Type_WellKnownType) (string, error) {
	switch w {
	case exprpb.Type_ANY:
		return "0", nil
	case exprpb.Type_TIMESTAMP:
		return "timestamp(0)", nil
	case exprpb.Type_DURATION:
		return "duration(0)", nil
	default:
		return "", fmt.Errorf("unknown well known type: %v", w)
	}
}

func zeroValue(tp *exprpb.Type) (string, error) {
	switch t := tp.TypeKind.(type) {
	case *exprpb.Type_Dyn:
		return "0", nil
	case *exprpb.Type_Null:
		return "null", nil
	case *exprpb.Type_Primitive:
		return zeroValuePrimitive(t.Primitive)
	case *exprpb.Type_Wrapper:
		return zeroValuePrimitive(t.Wrapper)
	case *exprpb.Type_WellKnown:
		return zeroValueWellKnown(t.WellKnown)
	case *exprpb.Type_ListType_:
		return "[]", nil
	case *exprpb.Type_MapType_:
		return "{}", nil
	case *exprpb.Type_Function:
		return "", fmt.Errorf("bad_type_function")
	case *exprpb.Type_MessageType:
		return t.MessageType + "{}", nil
	case *exprpb.Type_TypeParam:
		return "0", nil
	case *exprpb.Type_Type:
		return "type", nil
	case *exprpb.Type_Error:
		return "", fmt.Errorf("error type")
	case *exprpb.Type_AbstractType_:
		return "", fmt.Errorf("abstract type %s", t.AbstractType.Name)
	default:
		return "", fmt.Errorf("unknown type kind: %v", tp.GetTypeKind())
	}
}

func genParams(params []*exprpb.Type) ([]string, error) {
	var args []string
	for _, param := range params {
		arg, err := zeroValue(param)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}
	return args, nil
}

func isIdentifier(s string) bool {
	match, err := regexp.Match("^[_a-zA-Z][_a-zA-Z0-9]*$", []byte(s))
	return err == nil && match && s != "_in_"
}

var operators = []string {
	"!_", "-_", "_!=_", "_%_", "_&&_", "_*_", "_+_", "_-_", "_/_",
	"_<=_", "_<_", "_==_", "_>=_", "_>_", "_?_:_", "_[_]", "_||_",
}

func genExprOverload(name string, o *exprpb.Decl_FunctionDecl_Overload) (string, error) {
	if len(o.Params) == 0 {
		return name + "()", nil
	}
	args, err := genParams(o.Params)
	if err != nil {
		return "", err
	}
	if isIdentifier(name) {
		var prog string
		first := true
		if o.IsInstanceFunction {
			prog = args[0] + "." + name + "("
		} else {
			prog = name + "(" + args[0]
			first = false
		}
		for i := 1; i < len(args); i++ {
			if !first {
				prog += ", "
			}
			first = false
			prog += args[i]
		}
		prog += ")"
		return prog, nil
	}
	for _, op := range operators {
		if op == name {
			var prog string
			i := 0
			for _, c := range op {
				if c == '_' {
					if i >= len(args) {
						return "", fmt.Errorf("not enough params: %v", o)
					}
					prog += "(" + args[i] + ")"
					i++
				} else {
					prog += string(c)
				}
			}
			if i != len(args) {
				return "", fmt.Errorf("too many params: %v", o)
			}
			return prog, nil
		}
	}
	if name == "@in" {
		if len(args) != 2 {
			return "", fmt.Errorf("wrong number params: %v", o)
		}
		return args[0] + " in " + args[1], nil
	}
	return "", fmt.Errorf("function %s neither identifier nor operator", name)
}

func (r *runConfig) TestDecl(t *testing.T, decl *exprpb.Decl) {
	switch d := decl.DeclKind.(type) {
	case *exprpb.Decl_Ident:
		// For identifiers, the name itself is a suitable program.
		err := r.runProg(decl.Name, decl.Name, verifyIdentifier)
		if err != nil {
			t.Error(err)
		}
	case *exprpb.Decl_Function:
		for _, o := range d.Function.Overloads {
			t.Run(o.OverloadId, func (tt *testing.T) {
				prog, err := genExprOverload(decl.Name, o)
				if err != nil {
					tt.Fatal(err)
				}
				err = r.runProg(o.OverloadId, prog, verifyOverload)
				if err != nil {
					tt.Error(err)
				}
			})
		}
	default:
		t.Errorf("unknown decl kind %v", decl.DeclKind)
	}
}

func (r *runConfig) runProg(name, prog string, ok verifier) error {
	// Parse
	preq := exprpb.ParseRequest{
		CelSource:	prog,
		SourceLocation:	"test",
	}
	pres, err := r.parseClient.Parse(context.Background(), &preq)
	if err != nil {
		return fmt.Errorf("%s: Fatal parse RPC error: %s %v", name, prog, err)
	}
	if pres == nil {
		return fmt.Errorf("%s: Empty parse RPC response", name)
	}
	parsedExpr := pres.ParsedExpr
	if parsedExpr == nil {
		return fmt.Errorf("%s: Fatal parse errors: %s %v", name, prog, pres.Issues)
	}
	if parsedExpr.Expr == nil {
		return fmt.Errorf("%s: parse returned empty root expression", name)
	}

	// Skip the check phase - we're not testing the checker.

	// Eval
	ereq := exprpb.EvalRequest{
		ExprKind:  &exprpb.EvalRequest_ParsedExpr{ParsedExpr: parsedExpr},
	}
	eres, err := r.evalClient.Eval(context.Background(), &ereq)
	if err != nil {
		return fmt.Errorf("fatal eval RPC error for %s: %v", name, err)
	}
	if eres == nil || eres.Result == nil {
		return fmt.Errorf("empty eval response for %s", name)
	}
	if !ok(eres.Result) {
		return fmt.Errorf("unsupported: %s", name)
	}
	return nil
}
