// package main defines a simple implementation of the pipe protocol.
//
// This is intended for unit tests and only supports parse '1 + 1'
// and eval (which deterministically returns 2).
package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	confpb "google.golang.org/genproto/googleapis/api/expr/conformance/v1alpha1"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
)

type codec struct {
	useBase64 bool
}

func (c *codec) marshal(in proto.Message) (string, error) {
	if c.useBase64 {
		bytes, err := proto.Marshal(in)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(bytes), nil
	} else {
		return protojson.MarshalOptions{}.Format(in), nil
	}
}

func (c *codec) unmarshal(encoded string, out proto.Message) error {
	if c.useBase64 {
		protoBytes, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			return err
		}
		return proto.Unmarshal(protoBytes, out)
	} else {
		return protojson.Unmarshal([]byte(encoded), out)
	}
}

func (c *codec) serialize(w *bufio.Writer, out proto.Message) error {
	output, err := c.marshal(out)
	if err != nil {
		return err
	}
	_, err = w.WriteString(output + "\n")
	if err != nil {
		return err
	}
	return w.Flush()
}

func getExampleExpr() *exprpb.ParsedExpr {
	return &exprpb.ParsedExpr{
		Expr: &exprpb.Expr{
			ExprKind: &exprpb.Expr_CallExpr{
				CallExpr: &exprpb.Expr_Call{
					Function: "_+_",
					Args: []*exprpb.Expr{
						{
							ExprKind: &exprpb.Expr_ConstExpr{
								ConstExpr: &exprpb.Constant{
									ConstantKind: &exprpb.Constant_Int64Value{
										Int64Value: 1,
									},
								},
							},
						},
						{
							ExprKind: &exprpb.Expr_ConstExpr{
								ConstExpr: &exprpb.Constant{
									ConstantKind: &exprpb.Constant_Int64Value{
										Int64Value: 1,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func processLoop() int {
	useBase64 := flag.Bool("use_base64", false, "use_base64 sets the pipe format to base64")
	flag.Parse()
	c := codec{
		useBase64: *useBase64,
	}
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for {

		cmd, err := reader.ReadString('\n')
		if err != nil {
			return 1
		}

		cmd = strings.TrimSpace(cmd)
		if cmd == "exit" {
			return 0
		}
		msg, err := reader.ReadString('\n')
		if err != nil {
			return 1
		}

		switch cmd {
		case "parse":
			req := confpb.ParseRequest{}
			if err := c.unmarshal(msg, &req); err != nil {
				fmt.Fprintf(os.Stderr, "bad parse req: %v\n", err)
				return 1
			}
			resp := confpb.ParseResponse{
				Issues: []*statuspb.Status{
					{
						Code:    int32(codes.Unimplemented),
						Message: fmt.Sprintf("parse (%s)", req.CelSource),
					},
				},
			}
			if req.CelSource == "1 + 1" {
				resp = confpb.ParseResponse{
					ParsedExpr: getExampleExpr(),
				}
			}

			if err = c.serialize(writer, &resp); err != nil {
				fmt.Fprintf(os.Stderr, "error serializing parse resp %v\n", err)
				return 1
			}

		case "eval":
			req := confpb.EvalRequest{}
			if err := c.unmarshal(msg, &req); err != nil {
				fmt.Fprintf(os.Stderr, "bad eval req: %v\n", err)
				return 1
			}

			// toy example only works for 1 + 1
			resp := confpb.EvalResponse{
				Result: &exprpb.ExprValue{
					Kind: &exprpb.ExprValue_Value{
						Value: &exprpb.Value{
							Kind: &exprpb.Value_Int64Value{Int64Value: 2},
						},
					},
				},
			}
			if err = c.serialize(writer, &resp); err != nil {
				fmt.Fprintf(os.Stderr, "error serializing parse resp %v\n", err)
				return 1
			}
		case "check":
			req := confpb.CheckRequest{}
			if err := c.unmarshal(msg, &req); err != nil {
				fmt.Fprintf(os.Stderr, "bad check req: %v\n", err)
				return 1
			}
			resp := confpb.CheckResponse{
				Issues: []*statuspb.Status{
					{
						Code:    int32(codes.Unimplemented),
						Message: "check unimplemented",
					},
				},
			}
			if err = c.serialize(writer, &resp); err != nil {
				fmt.Fprintf(os.Stderr, "error serializing check resp %v\n", err)
				return 1
			}
		default:
			fmt.Fprintf(os.Stderr, "unsupported cmd: %s\n", cmd)
			return 1
		}
	}
}

func main() {
	rc := processLoop()
	os.Exit(rc)
}
