package main

import (
	"log"
	"os"

	"github.com/google/cel-spec/tests/simple"
)

func main() {
	os.Exit(mainHelper())
}

func mainHelper() int {
	err := setup()
	defer shutdown()
	if err != nil {
		log.Fatal(err)
		return 1
	}
	return simple.Run()
}


