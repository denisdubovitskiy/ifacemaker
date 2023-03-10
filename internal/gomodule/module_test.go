package gomodule

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/spf13/afero"

	"github.com/Masterminds/semver"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	cases := []string{
		"01_vault_v1_with_version.yml",
		"02_vault_v1_automatic_version.yml",
		"03_mattermost_v5_automatic_version.yml",
		"04_mattermost_v5_with_version.yml",
		"05_net_http.yml",
	}

	for _, tc := range cases {
		tcFile := tc

		t.Run(tcFile, func(t *testing.T) {
			t.Parallel()

			tc := parseTestCase(t, filepath.Join("testdata", tcFile))

			parser := newParser()
			parser.fs = afero.NewMemMapFs()
			parser.modcache = func() string { return tc.Mock.GOMODCACHE }
			for _, d := range tc.Mock.Dirs {
				_ = parser.fs.MkdirAll(filepath.Join(tc.Mock.GOMODCACHE, d), os.ModePerm) //nolint:errcheck
			}

			// act
			got, err := parser.Parse(tc.Given.Name, "")
			got.gomodcache = func() string { return tc.Mock.GOMODCACHE }
			got.goroot = func() string { return tc.Mock.GOROOT }

			gotDir := got.Directory("")
			gotVersionDir := got.VersionDirectory()

			// assert
			require.NoError(t, err)
			require.Equal(t, tc.Want.Module.Name, got.Name)
			require.Equal(t, tc.Want.Module.Base, got.Base)
			require.Equal(t, tc.Want.Module.Dir, got.Dir)
			require.Equal(t, tc.Want.Module.Sem, got.Ver)
			require.Equal(t, tc.Want.Directory, gotDir)
			require.Equal(t, tc.Want.VersionDirectory, gotVersionDir)
		})
	}
}

func TestSortVersions(t *testing.T) {
	cases := []struct {
		given []*semver.Version
		want  []*semver.Version
	}{
		{
			given: []*semver.Version{
				semver.MustParse("v5@v5.39.2"),
				semver.MustParse("v5@v5.39.3"),
			},
			want: []*semver.Version{
				semver.MustParse("v5@v5.39.3"),
				semver.MustParse("v5@v5.39.2"),
			},
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run("", func(t *testing.T) {
			t.Parallel()

			// act
			sortVersions(tc.given)

			// assert
			require.Equal(t, tc.want, tc.given)
		})
	}
}

type moduleTestGiven struct {
	Name string `yaml:"name"`
}

type moduleTestModule struct {
	Name string `yaml:"name"`
	Base string `yaml:"base"`
	Dir  string `yaml:"dir"`
	Ver  string `yaml:"ver"`
	Sem  *semver.Version
}

type moduleTestWant struct {
	Directory        string           `yaml:"directory"`
	VersionDirectory string           `yaml:"version_directory"`
	Module           moduleTestModule `yaml:"module"`
}

type moduleTestMock struct {
	GOMODCACHE string   `yaml:"gomodcache"`
	GOROOT     string   `yaml:"goroot"`
	Dirs       []string `yaml:"dirs"`
}

type moduleTestCase struct {
	Given moduleTestGiven `yaml:"given"`
	Want  moduleTestWant  `yaml:"want"`
	Mock  moduleTestMock  `yaml:"mock"`
}

func parseTestCase(t *testing.T, filename string) moduleTestCase {
	t.Helper()
	c, err := os.ReadFile(filename)
	require.NoError(t, err)

	var tc moduleTestCase
	err = yaml.Unmarshal(c, &tc)
	require.NoError(t, err)

	if len(tc.Want.Module.Ver) > 0 {
		tc.Want.Module.Sem = semver.MustParse(tc.Want.Module.Ver)
	}

	return tc
}
