// Package celrpc defines CEL conformance service RPC helpers.
package celrpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	confpb "google.golang.org/genproto/googleapis/api/expr/conformance/v1alpha1"
)

// ConfClient manages calls to conformance test services.
type ConfClient struct {
	confpb.ConformanceServiceClient
	cmd  *exec.Cmd
	conn *grpc.ClientConn
}

// NewClientFromPath creates a new ConformanceService gRPC client,
// connecting to a server which is launched by the binary at the given
// serverPath.  The spawned server shares the current process's stderr,
// so its log messages will be visible.  The caller must call Shutdown()
// on the retured ConfClient, even if NewClientFromPath() returns a
// non-nil error.
func NewClientFromPath(serverPath string) (*ConfClient, error) {
	c := ConfClient{}

	cmd := exec.Command(serverPath)
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

// ExampleNewClientFromPath creates a new CEL RPC client using a path to a server binary.
// TODO Run from celrpc_test.go.
func ExampleNewClientFromPath() {
	c, err := NewClientFromPath("/path/to/server/binary")
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

// Shutdown deallocates all resources associated with the client.
// No further calls should be made on the client after shutdown.
// Shutdown should be called even on an error return from NewClientFromPath().
func (c *ConfClient) Shutdown() {
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
