package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"

	"github.com/denisdubovitskiy/ifacemaker/internal/generator"
	"github.com/denisdubovitskiy/ifacemaker/internal/gomodule"
	"github.com/jessevdk/go-flags"
)

type arguments struct {
	SourcePackage  string `short:"s" long:"source-pkg" description:"Go import path to struct" required:"true"`
	SourceVersion  string `short:"v" long:"source-version" description:"Semantic version of the source package (example: v1.9.0)" required:"false"`
	ModulePath     string `short:"m" long:"module-path" description:"Submodule path from the root" required:"false"`
	ResultPackage  string `short:"p" long:"result-pkg" description:"Result package name" required:"true"`
	StructName     string `short:"t" long:"struct-name" description:"A structure name to generate interface for" required:"true"`
	InterfaceName  string `short:"i" long:"interface-name" description:"Name of the generated interface" required:"true"`
	OutputFileName string `short:"o" long:"output" description:"OutputFileName file name" required:"true"`
}

// ifacemaker \
// --source-pkg \
// github.com/mattermost/mattermost-server/v5@v5.39.3 \
// --module-path model \
// --result-pkg client \
// --struct-name Client4 \
// --interface-name Client4 \
// --output mattermost/client.go

func main() {
	var args arguments

	if _, err := flags.ParseArgs(&args, os.Args); err != nil {
		if flags.WroteHelp(err) {
			return
		}

		os.Exit(1)
	}

	module, err := gomodule.Parse(args.SourcePackage, args.SourceVersion)
	if err != nil {
		log.Fatal(err)
	}

	files, err := findSourceFiles(module.Directory(args.ModulePath))
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
	if err := os.MkdirAll(filepath.Dir(args.OutputFileName), os.ModePerm); err != nil {
		log.Fatal(err.Error())
	}
	if err := os.WriteFile(args.OutputFileName, generatedCode, 0644); err != nil {
		log.Fatal(err.Error())
	}
}

func newSourceFilesFinder() *sourceFilesFinder {
	return &sourceFilesFinder{fs: afero.NewOsFs()}
}

type sourceFilesFinder struct {
	fs afero.Fs
}

func (f *sourceFilesFinder) findSourceFiles(directory string) ([]string, error) {
	var files []string

	entries, err := afero.ReadDir(f.fs, directory)
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

var findSourceFiles = newSourceFilesFinder().findSourceFiles
