package lib

import (
	"io"
	"runtime"
	"strings"
)

var crlfNewline = runtime.GOOS == "windows"

func WriteOutput(w io.Writer, s string) error {
	if crlfNewline {
		// First normalize to \n, then convert all newlines to \r\n
		// (on Windows).
		s = strings.Replace(s, "\r\n", "\n", -1)
		s = strings.Replace(s, "\n", "\r\n", -1)
	}
	_, err := io.WriteString(w, s)
	return err
}
