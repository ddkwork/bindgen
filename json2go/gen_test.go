package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	os.RemoveAll("tmp")
	Walk("testdata")
}
