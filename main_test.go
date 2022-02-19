package main_test

import (
	"io"
	"os"
	"runtime"
)

func Example() {
	if runtime.GOOS == "windows" {
		io.WriteString(os.Stdout, "foo\r\nbar\r\n")
	} else {
		io.WriteString(os.Stdout, "foo\nbar\n")
	}
	// Output:
	// foo
	// bar
}
