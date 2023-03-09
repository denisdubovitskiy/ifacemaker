package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/denisdubovitskiy/ifacemaker/generator"
	"github.com/denisdubovitskiy/ifacemaker/gopath"
	"github.com/jessevdk/go-flags"
)

type arguments struct {
	SourcePackage  string `short:"s" long:"source-pkg" description:"Go import path to struct" required:"true"`
	ModulePath     string `short:"m" long:"module-path" description:"Submodule path from the root" required:"false"`
	ResultPackage  string `short:"p" long:"result-pkg" description:"Result package name" required:"true"`
	StructName     string `short:"t" long:"struct-name" description:"A structure name to generate interface for" required:"true"`
	InterfaceName  string `short:"i" long:"interface-name" description:"Name of the generated interface" required:"true"`
	OutputFileName string `short:"o" long:"output" description:"OutputFileName file name" required:"true"`
}

// --source-pkg github.com/mattermost/mattermost-server/v5 \
// --result-pkg mattermost \
// --struct-name Audit \
// --module-path model \
// --interface-name Audit \
// --output mattermost/audit.go

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

	gomod := filepath.Join(gopath.Find(), "pkg", "mod")

	module, err := parseModule(args.SourcePackage)
	if err != nil {
		log.Fatal(err)
	}

	version := module.Version
	if version == "" {
		directory := filepath.Join(gomod, module.Name)
		dirs, err := os.ReadDir(directory)
		if err != nil {
			log.Fatalf("trying to determine a last version, reading %s: %v", directory, err)
		}
		versions := encodeDirsToStrings(dirs)
		sortVersions(versions)
		version = versions[0]
	}

	directory := filepath.Join(gomod, module.Name, version, args.ModulePath)

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
	if err := os.MkdirAll(filepath.Dir(args.OutputFileName), os.ModePerm); err != nil {
		log.Fatal(err.Error())
	}
	if err := os.WriteFile(args.OutputFileName, generatedCode, 0644); err != nil {
		log.Fatal(err.Error())
	}
}

func encodeDirsToStrings(dirs []os.DirEntry) []string {
	result := make([]string, len(dirs))
	for i, d := range dirs {
		result[i] = d.Name()
	}
	return result
}

type sourcePackage struct {
	Name         string
	Version      string
	MajorVersion string
}

func parseModule(p string) (*sourcePackage, error) {
	version := ""
	module := p

	if strings.Contains(p, "@") {
		parts := strings.Split(p, "@")
		version = parts[1]
		if !strings.HasPrefix(version, "v") {
			return nil, fmt.Errorf("validation error: version should start with v")
		}

		module = parts[0]
	}

	majorVersion := ""

	// module has major version
	if lastSlash := strings.LastIndex(module, "/"); lastSlash > -1 {
		major := module[lastSlash+1:]
		matched, err := regexp.MatchString(`^v\d+$`, major)
		if err != nil {
			return nil, fmt.Errorf("major version check failed: %s - %v", major, err)
		}
		if matched {
			majorVersion = major
			module = module[:lastSlash]
		}
	}

	return &sourcePackage{
		Name:         module,
		Version:      version,
		MajorVersion: majorVersion,
	}, nil
}

type versionDirectory struct {
	major   int
	version *semver.Version
}

func parseVersionDirectory(v string) versionDirectory {
	var (
		major   int
		version string
	)

	if strings.Contains(v, "@") {
		atIdx := strings.Index(v, "@")
		version = v[atIdx+1:]
		majorStr := v[1:atIdx]
		major, _ = strconv.Atoi(majorStr)
	}

	sem := semver.MustParse(version)

	return versionDirectory{
		major:   major,
		version: sem,
	}
}

func sortVersions(versions []string) {
	sort.Slice(versions, func(i, j int) bool {
		return compareTwoVersions(parseVersionDirectory(versions[i]), parseVersionDirectory(versions[j]))
	})
}

func compareTwoVersions(v1, v2 versionDirectory) bool {
	if v1.major > v2.major {
		return true
	}

	return v1.version.GreaterThan(v2.version)
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
