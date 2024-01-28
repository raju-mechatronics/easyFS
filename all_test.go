package gofs

import (
	"fmt"
	"testing"
)

// test all package
func TestGoFS(t *testing.T) {
	fs := FS("/home/username")
	fmt.Println("printing the path", fs)
	if fs.String() != "/home/username" {
		t.Error("FS.String() failed")
	}
}
