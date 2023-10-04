package celrpc

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/bazelbuild/rules_go/go/tools/bazel"

	confpb "google.golang.org/genproto/googleapis/api/expr/conformance/v1alpha1"
)

var (
	serverCmd        string
	serverBase64Flag string
)

func init() {
	flag.StringVar(&serverCmd, "server_cmd", "", "cmd to start conformance service implementation")
	flag.StringVar(&serverBase64Flag, "server_base64_flag", "-use_base64", "flag the pipe server uses to enable base64 encoding")
	flag.Parse()
}

func TestPipeParse(t *testing.T) {
	for _, useBase64 := range []bool{false, true} {
		t := t
		t.Run(fmt.Sprintf("useBase64=%v", useBase64), func(t *testing.T) {
			serverCmd, err := bazel.Runfile(serverCmd)

			if useBase64 {
				serverCmd = fmt.Sprintf("%s %s", serverCmd, serverBase64Flag)
			}

			if err != nil {
				t.Fatalf("error loading bazel runfile path, %v", err)
			}
			conf, err := NewPipeClient(serverCmd, useBase64, false /*usePings*/)
			defer conf.Shutdown()
			if err != nil {
				t.Fatalf("error initializing client got %v wanted nil", err)
			}

			r := make(chan *confpb.ParseResponse)
			e := make(chan error)
			go func() {
				resp, err := conf.Parse(context.Background(), &confpb.ParseRequest{
					CelSource: "1 + 1",
				})
				e <- err
				r <- resp
			}()

			var resp *confpb.ParseResponse
			select {
			case <-time.After(2 * time.Second):
				err = errors.New("timeout")
			case err = <-e:
				resp = <-r
			}

			if err != nil {
				t.Fatalf("error from pipe: %v", err)
			}

			if len(resp.Issues) > 0 {
				t.Errorf("Issues: got %v expected none", resp.Issues)
			}

			if resp.GetParsedExpr().GetExpr().GetCallExpr().GetFunction() != "_+_" {
				t.Errorf("unexpected ast got: %s wanted _+_(1, 1)", resp.GetParsedExpr())
			}
		})
	}
}

func TestPipeCrashRecover(t *testing.T) {
	for _, useBase64 := range []bool{false, true} {
		t := t
		t.Run(fmt.Sprintf("useBase64=%v", useBase64), func(t *testing.T) {
			serverCmd, err := bazel.Runfile(serverCmd)

			if useBase64 {
				serverCmd = fmt.Sprintf("%s %s", serverCmd, serverBase64Flag)
			}

			if err != nil {
				t.Fatalf("error loading bazel runfile path, %v", err)
			}
			conf, err := NewPipeClient(serverCmd, useBase64, true /*usePings*/)
			defer conf.Shutdown()
			if err != nil {
				t.Fatalf("error initializing client got %v wanted nil", err)
			}
			var resp *confpb.ParseResponse
			r := make(chan *confpb.ParseResponse)
			e := make(chan error)
			go func() {
				resp, err := conf.Parse(context.Background(), &confpb.ParseRequest{
					CelSource: "test_crash",
				})
				e <- err
				r <- resp
			}()

			select {
			case <-time.After(2 * time.Second):
				err = errors.New("timeout")
			case err = <-e:
				resp = <-r
			}

			if err == nil {
				t.Fatalf("Expected error from pipe, got nil")
			}

			go func() {
				resp, err := conf.Parse(context.Background(), &confpb.ParseRequest{
					CelSource: "1 + 1",
				})
				e <- err
				r <- resp
			}()

			select {
			case <-time.After(2 * time.Second):
				err = errors.New("timeout")
			case err = <-e:
				resp = <-r
			}

			if err != nil {
				t.Fatalf("error from pipe: %v", err)
			}

			if len(resp.Issues) > 0 {
				t.Errorf("Issues: got %v expected none", resp.Issues)
			}

			if resp.GetParsedExpr().GetExpr().GetCallExpr().GetFunction() != "_+_" {
				t.Errorf("unexpected ast got: %s wanted _+_(1, 1)", resp.GetParsedExpr())
			}
		})
	}

}

func TestPipeEval(t *testing.T) {
	for _, useBase64 := range []bool{false, true} {
		t := t
		t.Run(fmt.Sprintf("useBase64=%v", useBase64), func(t *testing.T) {
			serverCmd, err := bazel.Runfile(serverCmd)

			if useBase64 {
				serverCmd = fmt.Sprintf("%s %s", serverCmd, serverBase64Flag)
			}

			if err != nil {
				t.Fatalf("error loading bazel runfile path, %v", err)
			}
			conf, err := NewPipeClient(serverCmd, useBase64, false /*usePings*/)
			defer conf.Shutdown()
			if err != nil {
				t.Fatalf("error initializing client got %v wanted nil", err)
			}

			r := make(chan *confpb.EvalResponse)
			e := make(chan error)
			go func() {
				resp, err := conf.Eval(context.Background(), &confpb.EvalRequest{})
				e <- err
				r <- resp
			}()

			var resp *confpb.EvalResponse
			select {
			case <-time.After(2 * time.Second):
				err = errors.New("timeout")
			case err = <-e:
				resp = <-r
			}

			if err != nil {
				t.Fatalf("error from pipe: %v", err)
			}

			if len(resp.Issues) > 0 {
				t.Errorf("Issues: got %v expected none", resp.Issues)
			}

			if resp.GetResult().GetValue().GetInt64Value() != 2 {
				t.Errorf("unexpected eval response got: %s wanted 2", resp)
			}
		})
	}
}

func TestPipeCheck(t *testing.T) {
	for _, useBase64 := range []bool{false, true} {
		t := t
		t.Run(fmt.Sprintf("useBase64=%v", useBase64), func(t *testing.T) {
			serverCmd, err := bazel.Runfile(serverCmd)

			if useBase64 {
				serverCmd = fmt.Sprintf("%s %s", serverCmd, serverBase64Flag)
			}

			if err != nil {
				t.Fatalf("error loading bazel runfile path, %v", err)
			}
			conf, err := NewPipeClient(serverCmd, useBase64, false /*usePings*/)
			defer conf.Shutdown()
			if err != nil {
				t.Fatalf("error initializing client got %v wanted nil", err)
			}

			r := make(chan *confpb.CheckResponse)
			e := make(chan error)
			go func() {
				resp, err := conf.Check(context.Background(), &confpb.CheckRequest{})
				e <- err
				r <- resp
			}()

			var resp *confpb.CheckResponse
			select {
			case <-time.After(2 * time.Second):
				err = errors.New("timeout")
			case err = <-e:
				resp = <-r
			}

			if err != nil {
				t.Fatalf("error from pipe: %v", err)
			}

			if len(resp.Issues) < 1 {
				t.Errorf("Issues: got %v expected Unimplemented Error", resp.Issues)
			}
		})
	}

}
