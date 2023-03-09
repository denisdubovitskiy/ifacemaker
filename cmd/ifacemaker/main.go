package main

import (
	"fmt"
	"github.com/denisdubovitskiy/ifacemaker/internal/generator"
	"github.com/denisdubovitskiy/ifacemaker/internal/golang"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
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

	module, err := parseModule(args.SourcePackage, args.SourceVersion)
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

type sourcePackage struct {
	Name string
	Base string
	Dir  string
	Sem  *semver.Version
}

func (p sourcePackage) HasMajor() bool {
	return p.Sem.Major() > 0
}

func (p sourcePackage) IsThirdParty() bool {
	return strings.Contains(p.Name, ".")
}

func (p sourcePackage) VersionDirectory() string {
	if p.Sem.Major() > 0 {
		return "v" + strconv.Itoa(int(p.Sem.Major())) + "@v" + p.Sem.String()
	}
	return p.Base + "@v" + p.Sem.String()
}

func (p sourcePackage) Directory(modulePath string) string {
	if !p.IsThirdParty() {
		return filepath.Join(golang.GOROOT(), "src", p.Name)
	}

	if p.HasMajor() {
		return filepath.Join(golang.GOMODCACHE(), p.Dir, p.VersionDirectory(), modulePath)
	}

	return filepath.Join(golang.GOMODCACHE(), p.Dir, p.VersionDirectory(), modulePath)
}

func parseModule(modulePath, versionStr string) (*sourcePackage, error) {
	// stdlib module
	if !strings.Contains(modulePath, ".") {
		return &sourcePackage{Name: modulePath}, nil
	}

	var version *semver.Version
	module := modulePath

	majorVersion := ""
	moduleDir := filepath.Dir(module)
	moduleBase := filepath.Base(module)

	matched, err := regexp.MatchString(`^v\d+$`, moduleBase)
	if err != nil {
		return nil, fmt.Errorf("major version check failed: %s - %v", moduleBase, err)
	}

	if matched {
		majorVersion = moduleBase
		module = moduleDir
		moduleBase = filepath.Base(module)
	}

	if versionStr == "" {
		if strings.Contains(modulePath, "@") {
			parts := strings.Split(modulePath, "@")
			versionStr = parts[1]
			if !strings.HasPrefix(versionStr, "v") {
				return nil, fmt.Errorf("validation error: version should start with v")
			}
			module = parts[0]
		}
	}

	if versionStr == "" {
		directory := filepath.Join(golang.GOMODCACHE(), moduleDir)
		dirs, err := os.ReadDir(directory)
		if err != nil {
			return nil, fmt.Errorf("trying to determine a last version, reading %s: %v", directory, err)
		}

		versions := make([]*semver.Version, 0, len(dirs))

		for _, dir := range dirs {
			if (len(majorVersion) > 0 && strings.HasPrefix(dir.Name(), majorVersion)) ||
				(len(majorVersion) == 0 && strings.HasPrefix(dir.Name(), moduleBase)) {

				v := dir.Name()
				v = strings.TrimPrefix(v, majorVersion)
				v = strings.TrimPrefix(v, moduleBase)
				v = strings.TrimPrefix(v, "@")

				versions = append(versions, semver.MustParse(v))
			}
		}

		if len(versions) == 0 {
			return nil, fmt.Errorf("unable to find files in %s while parsing module version", directory)
		}

		sortVersions(versions)
		version = versions[0]
	}

	return &sourcePackage{
		Name: module,
		Base: moduleBase,
		Dir:  moduleDir,
		Sem:  version,
	}, nil
}

func sortVersions(versions []*semver.Version) {
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].GreaterThan(versions[j])
	})
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
