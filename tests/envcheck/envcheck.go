// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package envcheck checks runtime support of declarations.
//
// A set of checker declarations is scanned to produce a set of CEL
// parse trees, each of which is then sent to the runtime.
//
// Identifier declarations are compiled to an expression of just that
// identifiers.  For instance, the "int" type identifier produces:
//
// 	int
//
// Function declarations are compiled to a separate expression for
// each overload.  The expression is an invocation of the overload with
// "zeroish" arguments of the appropriate type.  The zeroish arguments
// are:
//
// 	int		0
// 	uint		0u
// 	double		0.0
// 	bool		false
// 	string		""
// 	bytes		b""
// 	null_type	null
// 	type		type
// 	list<A>		[]
// 	map<A,B>	{}
// 	enum E		0
// 	message M	M{}
//
// For instance, the "_/_" function with overloads
//
// 	_/_: (int, int) -> int
// 	_/_: (uint, uint) -> uint
// 	_/_: (double, double) -> double
//
// compiles to the expressions
//
// 	(0)/(0)
// 	(0u)/(0u)
// 	(0.0)/(0.0)
//
// which are then evaluated.
//
// This test suite does not check that the overloads are implemented
// correctly, only that they are implemented at all.  The test will pass
// unless the expression evaluates (with no bindings) to any result or
// error other than "no_matching_overload".  For instance, the first
// two expressions for _/_ will generate division-by-zero errors, but
// this will pass the test.
//
package envcheck

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/cel-spec/tools/celrpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	confpb "google.golang.org/genproto/googleapis/api/expr/conformance/v1alpha1"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	spb "google.golang.org/genproto/googleapis/rpc/status"
)

// runConfig holds the client stub for the server for the runtime.
type runConfig struct {
	client celrpc.ConfClient
}

var (
	// noOverload holds the proto representation of the "no_matching_overload" error.
	noOverload = &exprpb.ExprValue{Kind: &exprpb.ExprValue_Error{
		Error: &exprpb.ErrorSet{Errors: []*spb.Status{
			status.New(codes.InvalidArgument, "no_matching_overload").Proto(),
		}},
	}}
)

// exprGen is an expression generator.  It accepts the next expression Id to use and returns
// an Expr proto and the next unused expression Id.
type exprGen func(int64) (*exprpb.Expr, int64)

// genExpr runs an expression generator with 0 as the first expression Id.
func genExpr(gen exprGen) *exprpb.Expr {
	e, _ := gen(int64(0))
	return e
}

// exprNil generates a nil pointer as its Expr, not consuming any expression Ids.
// It's useful for uniform handling of nil values when construction Exprs.
var exprNil exprGen = func(id int64) (*exprpb.Expr, int64) {
	return nil, id
}

// exprIdent generates an Ident expression.
func exprIdent(name string) exprGen {
	return func(id int64) (*exprpb.Expr, int64) {
		return &exprpb.Expr{
			Id: id,
			ExprKind: &exprpb.Expr_IdentExpr{
				IdentExpr: &exprpb.Expr_Ident{
					Name: name,
				},
			},
		}, id + 1
	}
}

// exprConst generates a Constant (literal) expression.
func exprConst(c *exprpb.Constant) exprGen {
	return func(id int64) (*exprpb.Expr, int64) {
		return &exprpb.Expr{
			Id:       id,
			ExprKind: &exprpb.Expr_ConstExpr{ConstExpr: c},
		}, id + 1
	}
}

// exprCall generates a Call expression with the given arguments.
func exprCall(f string, args ...exprGen) exprGen {
	return exprCallTarget(f, exprNil, args...)
}

// exprCallTarget generates a Call expression with the given target and arguments.
func exprCallTarget(f string, target exprGen, args ...exprGen) exprGen {
	return func(id int64) (*exprpb.Expr, int64) {
		tExp, id := target(id)
		var argExp []*exprpb.Expr
		for _, arg := range args {
			e, i := arg(id)
			argExp = append(argExp, e)
			id = i
		}
		return &exprpb.Expr{
			Id: id,
			ExprKind: &exprpb.Expr_CallExpr{
				CallExpr: &exprpb.Expr_Call{
					Target:   tExp,
					Function: f,
					Args:     argExp,
				},
			},
		}, id + 1
	}
}

// emptyList generates an empty CreateList expression.
var emptyList exprGen = func(id int64) (*exprpb.Expr, int64) {
	return &exprpb.Expr{
		Id: id,
		ExprKind: &exprpb.Expr_ListExpr{
			ListExpr: &exprpb.Expr_CreateList{},
		},
	}, id + 1
}

// exprEmptyStruct generates an empty CreateStruct expression.  The messageName may be nil,
// indicating an empty map.
func exprEmptyStruct(messageName string) exprGen {
	return func(id int64) (*exprpb.Expr, int64) {
		return &exprpb.Expr{
			Id: id,
			ExprKind: &exprpb.Expr_StructExpr{
				StructExpr: &exprpb.Expr_CreateStruct{
					MessageName: messageName,
				},
			},
		}, id + 1
	}
}

// emptyMap generates a CreateStruct expression for an empty map.
var emptyMap = exprEmptyStruct("")

// Generators for zero-ish constants.
var (
	zeroConstBool   = exprConst(&exprpb.Constant{ConstantKind: &exprpb.Constant_BoolValue{}})
	zeroConstInt    = exprConst(&exprpb.Constant{ConstantKind: &exprpb.Constant_Int64Value{}})
	zeroConstUint   = exprConst(&exprpb.Constant{ConstantKind: &exprpb.Constant_Uint64Value{}})
	zeroConstDouble = exprConst(&exprpb.Constant{ConstantKind: &exprpb.Constant_DoubleValue{}})
	zeroConstString = exprConst(&exprpb.Constant{ConstantKind: &exprpb.Constant_StringValue{}})
	zeroConstBytes  = exprConst(&exprpb.Constant{ConstantKind: &exprpb.Constant_BytesValue{}})
	zeroConstNull   = exprConst(&exprpb.Constant{ConstantKind: &exprpb.Constant_NullValue{}})
)

// zeroValuePrimitive returns a generator for the zero-ish value of a primitive type.
func zeroValuePrimitive(p exprpb.Type_PrimitiveType) (exprGen, error) {
	switch p {
	case exprpb.Type_BOOL:
		return zeroConstBool, nil
	case exprpb.Type_INT64:
		return zeroConstInt, nil
	case exprpb.Type_UINT64:
		return zeroConstUint, nil
	case exprpb.Type_DOUBLE:
		return zeroConstDouble, nil
	case exprpb.Type_STRING:
		return zeroConstString, nil
	case exprpb.Type_BYTES:
		return zeroConstBytes, nil
	default:
		return nil, fmt.Errorf("unknown primitive type: %v", p)
	}
}

// zeroValueWellKnown returns a generator for the zero-ish value of a well-known type.
func zeroValueWellKnown(w exprpb.Type_WellKnownType) (exprGen, error) {
	switch w {
	case exprpb.Type_ANY:
		return zeroConstInt, nil
	case exprpb.Type_TIMESTAMP:
		return exprCall("timestamp", zeroConstInt), nil
	case exprpb.Type_DURATION:
		return exprCall("duration", zeroConstInt), nil
	default:
		return nil, fmt.Errorf("unknown well known type: %v", w)
	}
}

// zeroValue returns a generator for the zero-ish value of a Type.
func zeroValue(tp *exprpb.Type) (exprGen, error) {
	switch t := tp.TypeKind.(type) {
	case *exprpb.Type_Dyn:
		return zeroConstInt, nil
	case *exprpb.Type_Null:
		return zeroConstNull, nil
	case *exprpb.Type_Primitive:
		return zeroValuePrimitive(t.Primitive)
	case *exprpb.Type_Wrapper:
		return zeroValuePrimitive(t.Wrapper)
	case *exprpb.Type_WellKnown:
		return zeroValueWellKnown(t.WellKnown)
	case *exprpb.Type_ListType_:
		return emptyList, nil
	case *exprpb.Type_MapType_:
		return emptyMap, nil
	case *exprpb.Type_Function:
		return nil, fmt.Errorf("bad_type_function")
	case *exprpb.Type_MessageType:
		return exprEmptyStruct(t.MessageType), nil
	case *exprpb.Type_TypeParam:
		return zeroConstInt, nil
	case *exprpb.Type_Type:
		return exprIdent("type"), nil
	case *exprpb.Type_Error:
		return nil, fmt.Errorf("error type")
	case *exprpb.Type_AbstractType_:
		return nil, fmt.Errorf("abstract type %s", t.AbstractType.Name)
	default:
		return nil, fmt.Errorf("unknown type kind: %v", tp.GetTypeKind())
	}
}

// overloadExpr returns the generator for a call to a given overload with zero-ish arguments.
func overloadExpr(name string, o *exprpb.Decl_FunctionDecl_Overload) (exprGen, error) {
	var args []exprGen
	for _, param := range o.Params {
		arg, err := zeroValue(param)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}
	if o.IsInstanceFunction && len(args) > 0 {
		return exprCallTarget(name, args[0], args[1:]...), nil
	}
	return exprCall(name, args...), nil
}

// TestDecl checks whether the runtime supports a given declaration by generating
// a test expression and running it.
func (r *runConfig) TestDecl(t *testing.T, decl *exprpb.Decl) {
	switch d := decl.DeclKind.(type) {
	case *exprpb.Decl_Ident:
		// For identifiers, the name itself is a suitable program.
		prog := genExpr(exprIdent(decl.Name))
		result, err := r.runProg(decl.Name, prog)
		if err != nil {
			t.Fatal(err)
		}
		// Any value is okay, any error is a problem.
		_, ok := result.Kind.(*exprpb.ExprValue_Value)
		if !ok {
			t.Errorf("got %v, want value", result)
		}

	case *exprpb.Decl_Function:
		for _, o := range d.Function.Overloads {
			t.Run(o.OverloadId, func(tt *testing.T) {
				g, err := overloadExpr(decl.Name, o)
				if err != nil {
					tt.Fatal(err)
				}
				result, err := r.runProg(o.OverloadId, genExpr(g))
				if err != nil {
					tt.Fatal(err)
				}
				// Only the "no_matching_overload" error is a problem.
				if proto.Equal(result, noOverload) {
					tt.Error("no matching overload")
				}
			})
		}
	default:
		t.Errorf("unknown decl kind %v", decl.DeclKind)
	}
}

// runProg evaluates a given expression on the runtime server and returns the result.
func (r *runConfig) runProg(name string, prog *exprpb.Expr) (*exprpb.ExprValue, error) {
	parsedExpr := &exprpb.ParsedExpr{
		Expr:       prog,
		SourceInfo: &exprpb.SourceInfo{},
	}
	ereq := confpb.EvalRequest{
		ExprKind: &confpb.EvalRequest_ParsedExpr{ParsedExpr: parsedExpr},
	}
	eres, err := r.client.Eval(context.Background(), &ereq)
	if err != nil {
		return nil, fmt.Errorf("fatal eval RPC error for %s: %v", name, err)
	}
	if eres == nil || eres.Result == nil {
		return nil, fmt.Errorf("empty eval response for %s", name)
	}
	return eres.Result, nil
}
