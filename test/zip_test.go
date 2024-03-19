package test

import (
	"github.com/oneliang/util-golang/file"
	"testing"
)

func TestZip(t *testing.T) {
	_ = file.Zip("a.zip", "/Users/oneliang/golang/githubWorkspace/util-golang/yaml/yaml.go")
}
