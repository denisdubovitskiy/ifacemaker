package generator

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

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
	wd, _ := os.Getwd()
	modcache := filepath.Join(wd, ".modcache")

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
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			spec := testReadFile(t, tc.directory, "case.yml")
			var test testCase
			testUnmarshalYaml(t, spec, &test)
			testGetPackage(t, test.Module, modcache)
			want := testReadFileString(t, tc.directory, "out.txt")

			// act
			got, err := Generate(Options{
				Files:             encodeFiles(test.Files, modcache),
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

func testReadFile(t *testing.T, directory, file string) []byte {
	t.Helper()
	content, err := os.ReadFile(filepath.Join("testdata", directory, file))
	require.NoError(t, err)
	return content
}

func testReadFileString(t *testing.T, directory, file string) string {
	t.Helper()
	return string(testReadFile(t, directory, file))
}

func testGetPackage(t *testing.T, module, modcache string) {
	t.Helper()
	cmd := exec.Command("go", "get", module)
	cmd.Env = append(os.Environ(), "GOMODCACHE="+modcache)
	out, err := cmd.CombinedOutput()
	require.NoErrorf(t, err, "cmd output: %s", string(out))
}

func testUnmarshalYaml(t *testing.T, in []byte, out interface{}) {
	t.Helper()
	err := yaml.Unmarshal(in, out)
	require.NoError(t, err)
}
