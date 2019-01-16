package simple

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/cel-spec/tools/celrpc"

	spb "github.com/google/cel-spec/proto/test/v1/simple"
)

var (
	flag_server_cmd string
	flag_parse_server_cmd string
	flag_check_server_cmd string
	flag_eval_server_cmd string
	rc *runConfig
)

func init () {
	flag.StringVar(&flag_server_cmd, "server", "", "path to binary for server when no phase-specific server defined")
	flag.StringVar(&flag_parse_server_cmd, "parse_server", "", "path to binary for parse server")
	flag.StringVar(&flag_check_server_cmd, "check_server", "", "path to binary for check server")
	flag.StringVar(&flag_eval_server_cmd, "eval_server", "", "path to binary for eval server")
}

// Server binaries specified by flags
func initRunConfig() (*runConfig, error) {
	// Find the server binary for each phase
	p_cmd := flag_server_cmd
	if flag_parse_server_cmd != "" {
		p_cmd = flag_parse_server_cmd
	}
	if p_cmd == "" {
		return nil, fmt.Errorf("no parse server defined")
	}

	c_cmd := flag_server_cmd
	if flag_check_server_cmd != "" {
		c_cmd = flag_check_server_cmd
	}
	if c_cmd == "" {
		return nil, fmt.Errorf("no check server defined")
	}

	e_cmd := flag_server_cmd
	if flag_eval_server_cmd != "" {
		e_cmd = flag_eval_server_cmd
	}
	if e_cmd == "" {
		return nil, fmt.Errorf("no eval server defined")
	}

	// Only launch each required binary once
	servers := make(map[string]*celrpc.ConfClient)
	servers[p_cmd] = nil
	servers[c_cmd] = nil
	servers[e_cmd] = nil
	for cmd, _ := range servers {
		cli, err := celrpc.NewClientFromPath(cmd)
		if err != nil {
			return nil, err
		}
		servers[cmd] = cli
	}

	var rc runConfig
	rc.parseClient = servers[p_cmd]
	rc.checkClient = servers[c_cmd]
	rc.evalClient = servers[e_cmd]
	return &rc, nil
}

// File path specified by flag
// TODO(jimlarson) use the utility filter to do text to binary
// proto conversion, since the C++ implementation understands Any
// messages.
func parseSimpleFile(filename string) (*spb.SimpleTestFile, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	s := string(bytes)
	var pb spb.SimpleTestFile
	err = proto.UnmarshalText(s, &pb)
	if err != nil {
		return nil, err
	}
	return &pb, nil
}

// Usage: --server=<path-to-binary> testfile1 ...
func TestMain(m *testing.M) {
	if len(os.Args) == 1 {
		// Special case for no args beyond Arg[0].  When being
		// run as a test target within cel-spec, there is no
		// conformance server to test.  Exit cleanly in order
		// to keep the tests green.
		os.Exit(0)
	}
	flag.Parse()
	var err error
	rc, err = initRunConfig()
	if err != nil {
		// Silly Go has no method in M to log errors or abort,
		// so we'll have to do it outside of the testing module.
		log.Fatal("Can't initialize test server", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestSimpleFile(t *testing.T) {
	for _, filename := range flag.Args() {
		testFile, err := parseSimpleFile(filename)
		if err != nil {
			t.Fatalf("Can't parse input file %v: %v", filename, err)
		}
		t.Logf("Running tests in file %v\n", testFile.Name)
		for _, section := range testFile.Section {
			t.Logf("Running tests in section %v\n", section.Name)
			for _, test := range section.Test {
				desc := fmt.Sprintf("%s/%s/%s", testFile.Name, section.Name, test.Name)
				t.Run(desc, func(t *testing.T) {
					err := rc.RunTest(test)
					if err != nil {
						t.Fatal(err)
					}
				})
			}
		}
	}
}
