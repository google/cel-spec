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

package envcheck

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/google/cel-spec/tools/celrpc"

	"google.golang.org/protobuf/encoding/prototext"

	envpb "github.com/google/cel-spec/proto/test/v1/testpb"
)

var (
	flagServerCmd string
	rc            *runConfig
)

func init() {
	flag.StringVar(&flagServerCmd, "server", "", "path to binary for eval-phase server")
}

// initRunConfig launches the server binaries specified by the flag.
func initRunConfig() (*runConfig, error) {
	// Find the server binary for each phase
	cmd := flagServerCmd
	if cmd == "" {
		return nil, fmt.Errorf("no server defined")
	}

	cli, err := celrpc.NewGrpcClient(cmd)
	if err != nil {
		return nil, err
	}

	return &runConfig{client: cli}, nil
}

// parseEnvFile parses the given file which should be the textproto of an Env message.
func parseEnvFile(filename string) (*envpb.Env, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var pb envpb.Env
	err = prototext.Unmarshal(bytes, &pb)
	if err != nil {
		return nil, err
	}
	return &pb, nil
}

// TestMain sets up the conformance test server.
// Usage: --server=<path-to-binary> testfile1 ...
func TestMain(m *testing.M) {
	flag.Parse()
	var err error
	// When flags are specified construct the run config object. When no
	// flags have been specified, the run config will be nil and the tests
	// will early return.
	if len(os.Args) > 1 {
		rc, err = initRunConfig()
	}
	if err != nil {
		// Silly Go has no method in M to log errors or abort,
		// so we'll have to do it outside of the testing module.
		log.Fatal("Can't initialize test server", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

// TestEnv runs the envcheck test files specified no the command line.
func TestEnv(t *testing.T) {
	// Special case to handle test invocation without args. See TestMain()
	// early return comment.
	if rc == nil {
		return
	}
	// Run the flag-configured tests.
	for _, filename := range flag.Args() {
		envFile, err := parseEnvFile(filename)
		if err != nil {
			t.Fatalf("Can't parse input file %v: %v", filename, err)
		}
		t.Logf("Running tests in file %v\n", envFile.Name)
		for _, decl := range envFile.Decl {
			t.Run(decl.Name, func(tt *testing.T) {
				rc.TestDecl(tt, decl)
			})
		}
	}
}
