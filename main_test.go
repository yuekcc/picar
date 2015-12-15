package main

import (
	"testing"
)

func Test_picar(t *testing.T) {
	prefix := "test"
	renameOnly := false
	dir := []string{"aa", "bbb", "ccc"}
	picar(prefix, renameOnly, dir...)
}
