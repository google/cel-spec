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

	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

type ConfClient struct {
	exprpb.ConformanceServiceClient
	cmd *exec.Cmd
	conn *grpc.ClientConn
}

// New creates a new ConformanceService gRPC client, connecting to a server
// which is launched by the binary at the given serverPath.
// The spawned server shares the current process's stderr,
// so its log messages will be visible.
// The caller must call Shutdown() on the retured ConfClient,
// even if NewClient() returns a non-nil error.
func NewClient(serverPath string) (*ConfClient, error) {
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

	c.ConformanceServiceClient = exprpb.NewConformanceServiceClient(conn)
	return &c, nil
}

func ExampleNewClient() bool {
	c, err := NewClient("/path/to/server/binary")
	defer c.Shutdown()
	if err != nil {
		log.Fatalf("")
		return false
	}
	parseRequest := exprpb.ParseRequest{
		CelSource: "1 + 1",
	}
	parseResponse, err := c.Parse(context.Background(), &parseRequest)
	return parseResponse != nil && err == nil
}

// Shutdown deallocates all resources associated with the client.
// No further calls should be made on the client after shutdown.
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

// StartServer listens on a dynamically-allocated port on the loopback
// network device, prints its address and port to stdout, then starts
// a gRPC server on the socket with the given service callbacks.
func StartServer(service exprpb.ConformanceServiceServer) {
        lis, err := net.Listen("tcp", "127.0.0.1:")
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }

        fmt.Printf("Listening on %v\n", lis.Addr().String())
        os.Stdout.Sync()

        s := grpc.NewServer()
        exprpb.RegisterConformanceServiceServer(s, service)
        reflection.Register(s)
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
}
