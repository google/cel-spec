package simple

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/cel-spec/tools/celrpc"

	spb "github.com/google/cel-spec/proto/test/v1/simple"
)

type stringArray []string

//String implements flag.Value.String()
func (i *stringArray) String() string {
	return strings.Join(*i, " ")
}

//Set implements flag.Value.Set()
func (i *stringArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	flagServerCmd      string
	flagParseServerCmd string
	flagCheckServerCmd string
	flagEvalServerCmd  string
	skipFlags          stringArray
	rc                 *runConfig
)

func init() {
	flag.StringVar(&flagServerCmd, "server", "", "path to binary for server when no phase-specific server defined")
	flag.StringVar(&flagParseServerCmd, "parse_server", "", "path to binary for parse server")
	flag.StringVar(&flagCheckServerCmd, "check_server", "", "path to binary for check server")
	flag.StringVar(&flagEvalServerCmd, "eval_server", "", "path to binary for eval server")
	flag.Var(&skipFlags, "skip_test", "name(s) of tests to skip. can be set multiple times. to skip the following tests: f1/s1/t1, f1/s1/t2, f1/s2/*, f2/s3/t3, you give the arguments --skip_test=f1/s1/t1,t2;s2 --skip_test=f2/s3/t3")
	flag.Parse()
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

	cCmd := flagServerCmd
	if flagCheckServerCmd != "" {
		cCmd = flagCheckServerCmd
	}
	if cCmd == "" {
		return nil, fmt.Errorf("no check server defined")
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
	servers[cCmd] = nil
	servers[eCmd] = nil
	for cmd := range servers {
		cli, err := celrpc.NewClientFromPath(cmd)
		if err != nil {
			return nil, err
		}
		servers[cmd] = cli
	}

	var rc runConfig
	rc.parseClient = servers[pCmd]
	rc.checkClient = servers[cCmd]
	rc.evalClient = servers[eCmd]
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

func TestSimpleFile(t *testing.T) {
	// Special case to handle test invocation without args. See TestMain()
	// early return comment.
	if rc == nil {
		return
	}
	var skipTests []string
	for _, flagVal := range skipFlags {
		fileInd := strings.Index(flagVal, "/")
		splitFile := strings.SplitN(flagVal, "/", 2)
		fileName := splitFile[0]
		sectionString := splitFile[1]
		if fileInd < 1 || sectionString == "" {
			log.Fatal("skip_test argument must contain at least <file>/<section>, received ", flagVal)
		}
		for _, sectionVal := range strings.Split(sectionString, ";") {
			sections := strings.Count(sectionVal, "/")
			if sections > 1 {
				log.Fatal("Unable to parse skip_test flag for ", sectionVal)
			}
			if sections == 0 {
				if sectionVal == "" {
					log.Fatal("Empty string where should be section name")
				}
				 skipTests = append(skipTests, fileName+"/"+sectionVal)
			} else if sections == 1 {
				splitSection := strings.SplitN(sectionVal, "/", 2)
				sectionName := splitSection[0]
				testString := splitSection[1]
				if testString == "" {
					log.Fatal("Empty string where should be test name")
				}
				tests := strings.Split(testString, ",")
				for _, test := range tests {
					if test == "" {
						log.Fatal("Empty string where should be test name")
					}
					skipTests = append(skipTests, fileName+"/"+sectionName+"/"+test)
				}
			}
		}
	}
	// Run the flag-configured tests.
	for _, filename := range flag.Args() {
		testFile, err := parseSimpleFile(filename)
		if err != nil {
			t.Fatalf("Can't parse input file %v: %v", filename, err)
		}
		t.Logf("Running tests in file %v\n", testFile.Name)
		for _, section := range testFile.Section {
			sectionPath := testFile.Name + "/" + section.Name
			if contains(skipTests, sectionPath) {
				t.Logf("Skipping all tests in section %v\n", section.Name)
				continue
			}
			t.Logf("Running tests in section %v\n", section.Name)
			for _, test := range section.Test {
				testPath := sectionPath + "/" + test.Name
				if contains(skipTests, testPath) {
					t.Logf("Skipping test name %v\n", test.Name)
					continue
				}
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

// Strings can be stored in a sorted order to speed up checks if needed
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
