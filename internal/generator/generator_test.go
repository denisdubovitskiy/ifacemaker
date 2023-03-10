package generator

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/denisdubovitskiy/ifacemaker/internal/golang"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

type testCase struct {
	Module         string   `yaml:"module"`
	Files          []string `yaml:"files"`
	StructName     string   `yaml:"struct_name"`
	InterfaceName  string   `yaml:"interface_name"`
	OutPackageName string   `yaml:"out_package_name"`
}

func TestGenerate(t *testing.T) {
	cases := []struct {
		name      string
		directory string
	}{
		{
			name:      "mattermost Audit struct",
			directory: "01_audit",
		},
		{
			name:      "mattermost Audit struct rename",
			directory: "02_audit_rename",
		},
		{
			name:      "mattermost Audit package rename",
			directory: "03_audit_package_rename",
		},
		{
			name:      "mattermost Client4 struct",
			directory: "04_client4",
		},
		{
			name:      "mattermost User struct",
			directory: "05_user",
		},
	}

	modpath := filepath.Join(golang.GOPATH(), "pkg", "mod")

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			content, err := os.ReadFile(filepath.Join("testdata", tc.directory, "case.yml"))
			require.NoError(t, err)

			var test testCase
			err = yaml.Unmarshal(content, &test)
			require.NoError(t, err)

			cmd := exec.Command("go", "get", test.Module)
			cmd.Env = append(os.Environ(), "GOPATH="+golang.GOPATH())
			out, err := cmd.CombinedOutput()
			require.NoErrorf(t, err, "cmd output: %s", string(out))

			content, err = os.ReadFile(filepath.Join("testdata", tc.directory, "out.txt"))
			require.NoError(t, err)

			want := string(content)

			// act
			got, err := Generate(Options{
				Files:             encodeFiles(test.Files, modpath),
				StructName:        test.StructName,
				InterfaceName:     test.InterfaceName,
				OutputPackageName: test.OutPackageName,
			})

			// assert
			require.NoError(t, err)
			require.Equal(t, want, string(got))
		})
	}
}

func encodeFiles(files []string, modpath string) []string {
	result := make([]string, len(files))
	for i, f := range files {
		result[i] = filepath.Join(modpath, f)
	}
	return result
}
