package gopath

import (
	"go/build"
	"os"
)

func Find() string {
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		return gopath
	}

	return build.Default.GOPATH
}
