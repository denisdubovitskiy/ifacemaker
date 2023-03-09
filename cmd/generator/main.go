package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jessevdk/go-flags"

	"github.com/densdubovitskiy/ifacemaker/generator"
)

type arguments struct {
	SourcePackage  string `short:"s" long:"source-pkg" description:"Go import path to struct" required:"true"`
	ResultPackage  string `short:"p" long:"result-pkg" description:"Result package name" required:"true"`
	StructName     string `short:"t" long:"struct-name" description:"A structure name to generate interface for" required:"true"`
	InterfaceName  string `short:"i" long:"interface-name" description:"Name of the generated interface" required:"true"`
	OutputFileName string `short:"o" long:"output" description:"OutputFileName file name" required:"true"`
}

// --source-pkg github.com/hashicorp/vault@v1.8.2/api.Client \
// --result-pkg vault \
// --struct-name Client \
// --interface-name Client \
// --output result/vault/client.go
func main() {
	var args arguments

	if _, err := flags.ParseArgs(&args, os.Args); err != nil {
		if flags.WroteHelp(err) {
			return
		}

		os.Exit(1)
	}

	gopath := parseGopath()
	gomod := filepath.Join(gopath, "pkg", "mod")

	module, err := parsePackage(args.SourcePackage)
	if err != nil {
		log.Fatal(err)
	}

	directory := filepath.Join(gomod, module.Directory())

	files, err := findSourceFiles(directory)
	if err != nil {
		log.Fatal(err)
	}

	generatedCode, err := generator.Generate(generator.Options{
		Files:             files,
		StructName:        args.StructName,
		OutputPackageName: args.ResultPackage,
		InterfaceName:     args.InterfaceName,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := os.MkdirAll(filepath.Dir(args.OutputFileName), 0644); err != nil {
		log.Fatal(err.Error())
	}
	if err := os.WriteFile(args.OutputFileName, generatedCode, 0644); err != nil {
		log.Fatal(err.Error())
	}
}

type sourcePackage struct {
	Name    string
	Version string
	Struct  string
	Path    string
}

func (s *sourcePackage) Directory() string {
	return filepath.Join(s.Name, s.Path+"@"+s.Version)
}

func parsePackage(p string) (*sourcePackage, error) {
	if !strings.Contains(p, "@") {
		return nil, fmt.Errorf("validation error: struct spec should contain @")
	}

	if strings.Count(p, "@") > 1 {
		return nil, fmt.Errorf("validation error: struct spec should contain single @")
	}

	atIndex := strings.LastIndex(p, "@")
	module := p[:atIndex]

	rightPart := p[atIndex+1:]

	if !strings.HasPrefix(rightPart, "v") {
		return nil, fmt.Errorf("validation error: version should start with v")
	}

	firstSlashIdx := strings.Index(rightPart, "/")
	if firstSlashIdx < 0 {
		return nil, fmt.Errorf("validation error: there must")
	}

	version := rightPart[:firstSlashIdx]

	rightPart = rightPart[firstSlashIdx+1:]
	lastDotIdx := strings.LastIndex(rightPart, ".")
	structName := rightPart[lastDotIdx+1:]

	return &sourcePackage{
		Name:    module,
		Struct:  structName,
		Version: version,
		Path:    rightPart[:lastDotIdx],
	}, nil
}

func parseGopath() string {
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		return gopath
	}

	return build.Default.GOPATH
}

func findSourceFiles(directory string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		if strings.HasSuffix(e.Name(), "_test.go") ||
			!strings.HasSuffix(e.Name(), ".go") {
			continue
		}

		files = append(files, filepath.Join(directory, e.Name()))
	}

	return files, nil
}
