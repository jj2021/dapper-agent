package tools

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	expected := fmt.Sprintf("Hello\nthere\n")
	result, err := read("/home/james/github.com/dandy-dapper/pkg/tools/testdata/file.txt")

	if err != nil {
		t.Errorf("Read error: %s", err.Error())
		return
	}
	if fmt.Sprintf("%s", result) != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}
}
