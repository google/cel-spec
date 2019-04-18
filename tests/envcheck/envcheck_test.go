package envcheck

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/cel-spec/tools/celrpc"

	envpb "github.com/google/cel-spec/proto/test/v1/envcheck"
)

var (
	flagServerCmd      string
	flagParseServerCmd string
	flagEvalServerCmd  string
	rc                 *runConfig
)

func init() {
	flag.StringVar(&flagServerCmd, "server", "", "path to binary for server when no phase-specific server defined")
	flag.StringVar(&flagParseServerCmd, "parse_server", "", "path to binary for parse server")
	flag.StringVar(&flagEvalServerCmd, "eval_server", "", "path to binary for eval server")
}

// Server binaries specified by flags
func initRunConfig() (*runConfig, error) {
	// Find the server binary for each phase
	pCmd := flagServerCmd
	if flagParseServerCmd != "" {
		pCmd = flagParseServerCmd
	}
	if pCmd == "" {
		return nil, fmt.Errorf("no parse server defined")
	}

	eCmd := flagServerCmd
	if flagEvalServerCmd != "" {
		eCmd = flagEvalServerCmd
	}
	if eCmd == "" {
		return nil, fmt.Errorf("no eval server defined")
	}

	// Only launch each required binary once
	servers := make(map[string]*celrpc.ConfClient)
	servers[pCmd] = nil
	servers[eCmd] = nil
	for cmd, _ := range servers {
		cli, err := celrpc.NewClientFromPath(cmd)
		if err != nil {
			return nil, err
		}
		servers[cmd] = cli
	}

	var rc runConfig
	rc.parseClient = servers[pCmd]
	rc.evalClient = servers[eCmd]
	return &rc, nil
}

// File path specified by flag
func parseEnvFile(filename string) (*envpb.Env, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	s := string(bytes)
	var pb envpb.Env
	err = proto.UnmarshalText(s, &pb)
	if err != nil {
		return nil, err
	}
	return &pb, nil
}

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
			t.Run(decl.Name, func (tt *testing.T) {
				rc.TestDecl(tt, decl)
			})
		}
	}
}
