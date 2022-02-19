package main_test

import (
	"fmt"
	"os"

	"github.com/benhoyt/go-windows-examples/lib"
)

func Example1() {
	err := lib.WriteOutput(os.Stdout, "foo\nbar\n")
	if err != nil {
		panic(err)
	}
	// Output:
	// foo
	// bar
}

func Example2() {
	fmt.Print("foo\nbar\n")
	// Output:
	// foo
	// bar
}

func Example3() {
	fmt.Print("foo\r\nbar\r\n")
	// Output:
	// foo
	// bar
}
