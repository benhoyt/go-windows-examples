package main_test

import (
	"os"

	"github.com/benhoyt/go-windows-examples/lib"
)

func Example() {
	err := lib.WriteOutput(os.Stdout, "foo\nbar\n")
	if err != nil {
		panic(err)
	}
	// Output:
	// foo
	// bar
}
