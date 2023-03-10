package golang

import (
	"go/build"
	"os"
	"path/filepath"
)

func GOPATH() string {
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		return gopath
	}

	return build.Default.GOPATH
}

func GOROOT() string {
	if goroot := os.Getenv("GOROOT"); goroot != "" {
		return goroot
	}

	return build.Default.GOROOT
}

func GOMODCACHE() string {
	return filepath.Join(GOPATH(), "pkg", "mod")
}
