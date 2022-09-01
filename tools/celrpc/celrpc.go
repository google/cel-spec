// Package celrpc defines CEL conformance service RPC helpers.
package celrpc

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	confpb "google.golang.org/genproto/googleapis/api/expr/conformance/v1alpha1"
)

// ConfClient manages calls to conformance test services.
type ConfClient interface {
	confpb.ConformanceServiceClient
	// Shutdown deallocates all resources associated with the client.
	// No further calls should be made on the client after shutdown.
	// Shutdown should be called even on an error return.
	Shutdown()
}

// gRPC conformance service client
type grpcConfClient struct {
	confpb.ConformanceServiceClient
	cmd  *exec.Cmd
	conn *grpc.ClientConn
}

// pipe conformance client uses the following protocol:
//   * two lines are sent over input
//   * first input line is "parse", "check", or "eval"
//   * second input line is JSON of the corresponding request
//   * one output line is expected, repeat again.
type pipeConfClient struct {
	cmd       *exec.Cmd
	stdOut    *bufio.Reader
	stdIn     io.Writer
	useBase64 bool
}

// NewGrpcClient creates a new gRPC ConformanceService client. A server binary
// is launched given the command line serverCmd. The spawned server shares the
// current process's stderr, so its log messages will be visible.
// The caller must call Shutdown() on the retured ConfClient, even if
// NewGrpcClient() returns a non-nil error.
func NewGrpcClient(serverCmd string) (ConfClient, error) {
	c := grpcConfClient{}

	fields := strings.Fields(serverCmd)
	cmd := exec.Command(fields[0], fields[1:]...)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return &c, err
	}
	cmd.Stderr = os.Stderr // share our error stream

	err = cmd.Start()
	if err != nil {
		return &c, err
	}
	// Only assign cmd for stopping if it has successfully started.
	c.cmd = cmd

	// Expect a port only with gRPC
	var addr string
	_, err = fmt.Fscanf(out, "Listening on %s\n", &addr)
	out.Close()
	if err != nil {
		return &c, err
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return &c, err
	}
	c.conn = conn
	c.ConformanceServiceClient = confpb.NewConformanceServiceClient(conn)
	return &c, nil
}

// ExampleNewGrpcClient creates a new CEL RPC client using a path to a server binary.
// TODO Run from celrpc_test.go.
func ExampleNewGrpcClient() {
	c, err := NewGrpcClient("/path/to/server/binary")
	defer c.Shutdown()
	if err != nil {
		log.Fatal("Couldn't create client")
	}
	parseRequest := confpb.ParseRequest{
		CelSource: "1 + 1",
	}
	parseResponse, err := c.Parse(context.Background(), &parseRequest)
	if err != nil {
		log.Fatal("Couldn't parse")
	}
	parsedExpr := parseResponse.ParsedExpr
	evalRequest := confpb.EvalRequest{
		ExprKind: &confpb.EvalRequest_ParsedExpr{ParsedExpr: parsedExpr},
	}
	evalResponse, err := c.Eval(context.Background(), &evalRequest)
	if err != nil {
		log.Fatal("Couldn't eval")
	}
	fmt.Printf("1 + 1 is %v\n", evalResponse.Result.GetValue().GetInt64Value())
}

// NewPipeClient launches a server binary using the provided serverCmd
// command line. The spawned server shares the current process's stderr, so its
// log messages will be visible. stdin and stdout are used for communication.
// The caller must call Shutdown() on the retured ConfClient, even if the
// method returns a non-nil error.
//
// base64Encode enables base64Encoded messages (b64encode(Any.serializeToString))
func NewPipeClient(serverCmd string, base64Encode bool) (ConfClient, error) {
	c := pipeConfClient{
		useBase64: base64Encode,
	}

	fields := strings.Fields(serverCmd)
	if len(fields) < 1 {
		return &c, fmt.Errorf("server cmd '%s' invalid", serverCmd)
	}
	cmd := exec.Command(fields[0], fields[1:]...)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return &c, err
	}
	c.stdIn, err = cmd.StdinPipe()
	if err != nil {
		return &c, err
	}
	cmd.Stderr = os.Stderr // share our error stream

	err = cmd.Start()
	if err != nil {
		return &c, err
	}
	// Only assign cmd for stopping if it has successfully started.
	c.cmd = cmd
	c.stdOut = bufio.NewReader(out)
	return &c, nil
}

// ExampleNewPipeClient creates a new CEL pipe client using a path to a server binary.
// TODO Run from celrpc_test.go.
func ExampleNewPipeClient() {
	c, err := NewPipeClient("/path/to/server/binary", false)
	defer c.Shutdown()
	if err != nil {
		log.Fatal("Couldn't create client")
	}
	parseRequest := confpb.ParseRequest{
		CelSource: "1 + 1",
	}
	parseResponse, err := c.Parse(context.Background(), &parseRequest)
	if err != nil {
		log.Fatal("Couldn't parse")
	}
	parsedExpr := parseResponse.ParsedExpr
	evalRequest := confpb.EvalRequest{
		ExprKind: &confpb.EvalRequest_ParsedExpr{ParsedExpr: parsedExpr},
	}
	evalResponse, err := c.Eval(context.Background(), &evalRequest)
	if err != nil {
		log.Fatal("Couldn't eval")
	}
	fmt.Printf("1 + 1 is %v\n", evalResponse.Result.GetValue().GetInt64Value())
}

func (c *pipeConfClient) marshal(in proto.Message) (string, error) {
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

func (c *pipeConfClient) unmarshal(encoded string, out proto.Message) error {
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

func (c *pipeConfClient) pipeCommand(cmd string, in proto.Message, out proto.Message) error {
	if _, err := c.stdIn.Write([]byte(cmd + "\n")); err != nil {
		return err
	}

	pipeInput, err := c.marshal(in)
	if err != nil {
		return err
	}
	if _, err := c.stdIn.Write([]byte(pipeInput + "\n")); err != nil {
		return err
	}
	pipeOutput, err := c.stdOut.ReadString('\n')
	if err != nil {
		return err
	}
	return c.unmarshal(pipeOutput, out)
}

// Parse implements a gRPC client stub with both pipe and gRPC
func (c *pipeConfClient) Parse(ctx context.Context, in *confpb.ParseRequest, opts ...grpc.CallOption) (*confpb.ParseResponse, error) {
	out := &confpb.ParseResponse{}
	err := c.pipeCommand("parse", in, out)
	return out, err
}

// Check implements a gRPC client stub with both pipe and gRPC
func (c *pipeConfClient) Check(ctx context.Context, in *confpb.CheckRequest, opts ...grpc.CallOption) (*confpb.CheckResponse, error) {
	out := &confpb.CheckResponse{}
	err := c.pipeCommand("check", in, out)
	return out, err
}

// Eval implements a gRPC client stub with both pipe and gRPC
func (c *pipeConfClient) Eval(ctx context.Context, in *confpb.EvalRequest, opts ...grpc.CallOption) (*confpb.EvalResponse, error) {
	out := &confpb.EvalResponse{}
	err := c.pipeCommand("eval", in, out)
	return out, err
}

// Shutdown implements the interface stub.
func (c *pipeConfClient) Shutdown() {
	if c.cmd != nil {
		c.cmd.Process.Kill()
		c.cmd.Wait()
		c.cmd = nil
	}
}

// Shutdown implements the interface stub.
func (c *grpcConfClient) Shutdown() {
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	if c.cmd != nil {
		c.cmd.Process.Kill()
		c.cmd.Wait()
		c.cmd = nil
	}
}

// RunServer listens on a dynamically-allocated port on the loopback
// network device, prints its address and port to stdout, then starts
// a gRPC server on the socket with the given service callbacks.
// Note that this call doesn't return until ther server exits.
func RunServer(service confpb.ConformanceServiceServer) {
	lis, err := net.Listen("tcp4", "127.0.0.1:")
	if err != nil {
		lis, err = net.Listen("tcp6", "[::1]:0")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}

	// Must print to stdout, so the client can find the port.
	// So, no, this must be 'fmt', not 'log'.
	fmt.Printf("Listening on %v\n", lis.Addr())
	os.Stdout.Sync()

	s := grpc.NewServer()
	confpb.RegisterConformanceServiceServer(s, service)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
