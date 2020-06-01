package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCacheDirName(t *testing.T) {
	basePath, _ := filepath.Abs("./")
	t.Log(basePath)

	basePath, _ = filepath.Abs(os.Args[0])
	basePath = filepath.Dir(basePath)
	t.Log(basePath)
}
