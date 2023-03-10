package gomodule

import (
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/denisdubovitskiy/ifacemaker/internal/golang"
	"github.com/spf13/afero"
)

type Module struct {
	Name string
	Base string
	Dir  string
	Ver  *semver.Version

	// mocked in tests to be reproducible
	gomodcache func() string
	goroot     func() string
}

func (p Module) HasMajor() bool {
	return p.Ver.Major() > 1
}

func (p Module) IsThirdParty() bool {
	return strings.Contains(p.Name, ".")
}

func (p Module) VersionDirectory() string {
	if p.Ver == nil {
		return ""
	}

	if p.HasMajor() {
		return "v" + strconv.Itoa(int(p.Ver.Major())) + "@v" + p.Ver.String()
	}
	return p.Base + "@v" + p.Ver.String()
}

func (p Module) Directory(modulePath string) string {
	if !p.IsThirdParty() {
		return filepath.Join(p.goroot(), "src", p.Name)
	}

	if p.HasMajor() {
		return filepath.Join(p.gomodcache(), p.Dir, p.VersionDirectory(), modulePath)
	}

	return filepath.Join(p.gomodcache(), p.Dir, p.VersionDirectory(), modulePath)
}

type parser struct {
	// mocked in tests to be reproducible
	fs       afero.Fs
	modcache func() string
}

func newParser() *parser {
	return &parser{
		fs:       afero.NewOsFs(),
		modcache: golang.GOMODCACHE,
	}
}

func (p *parser) Parse(modulePath, versionStr string) (*Module, error) {
	// net/http
	if !strings.Contains(modulePath, ".") {
		return &Module{Name: modulePath}, nil
	}

	// github.com/mattermost/mattermost-server/v5@v5.39.3
	if versionStr == "" && strings.Contains(modulePath, "@") {
		parts := strings.Split(modulePath, "@")
		versionStr = parts[1]
		if !strings.HasPrefix(versionStr, "v") {
			return nil, fmt.Errorf("validation error: version should start with v")
		}
		modulePath = parts[0]
	}

	var version *semver.Version
	module := modulePath

	majorVersion := ""
	moduleDir := filepath.Dir(module)
	moduleBase := filepath.Base(module)

	matched, _ := regexp.MatchString(`^v\d+$`, moduleBase) //nolint:errcheck // compile error
	if matched {
		majorVersion = moduleBase
		module = moduleDir
		moduleBase = filepath.Base(module)
	}

	if versionStr == "" {
		directory := filepath.Join(p.modcache(), moduleDir)
		dirs, err := afero.ReadDir(p.fs, directory)
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
	} else {
		version = semver.MustParse(versionStr)
	}

	return &Module{
		Name: module,
		Base: moduleBase,
		Dir:  moduleDir,
		Ver:  version,

		goroot:     golang.GOROOT,
		gomodcache: golang.GOMODCACHE,
	}, nil
}

func sortVersions(versions []*semver.Version) {
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].GreaterThan(versions[j])
	})
}

var Parse = newParser().Parse
