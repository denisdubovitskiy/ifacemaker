package e2e

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
	ModulePath     string   `yaml:"module_path"`
}

func TestBinary(t *testing.T) {
	wd, _ := os.Getwd()
	modcache := filepath.Join(wd, ".modcache")

	path := "../cmd/ifacemaker/main.go"
	binary := filepath.Join(wd, "ifacemaker")

	// 1. build binary
	cmd := exec.Command("go", "build", "-o", binary, path)
	_, err := cmd.CombinedOutput()
	require.NoErrorf(t, err, "unable to build binary")

	// 2. defer cleanup binary
	t.Cleanup(func() {
		err := os.Remove(binary)
		require.NoErrorf(t, err, "unable to remove binary")
	})

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

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			spec := testReadFile(t, filepath.Join("testdata", tc.directory, "case.yml"))
			var test testCase
			testUnmarshalYaml(t, spec, &test)
			testGetPackage(t, test.Module, modcache)
			want := testReadFileString(t, filepath.Join("testdata", tc.directory, "out.txt"))

			cmd := exec.Command(binary,
				"--source-pkg", test.Module,
				"--module-path", test.ModulePath,
				"--result-pkg", test.OutPackageName,
				"--struct-name", test.StructName,
				"--interface-name", test.InterfaceName,
				"--output", tc.directory+".txt",
			)
			cmd.Env = append(os.Environ(), "GOMODCACHE="+modcache)
			out, err := cmd.CombinedOutput()
			require.NoErrorf(t, err, "cmd output: %s", string(out))
			t.Cleanup(func() {
				err := os.Remove(tc.directory + ".txt")
				require.NoError(t, err)
			})

			got := testReadFileString(t, tc.directory+".txt")

			// assert
			require.NoError(t, err)
			require.Equal(t, want, got)
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

func testReadFile(t *testing.T, file string) []byte {
	t.Helper()
	content, err := os.ReadFile(file)
	require.NoError(t, err)
	return content
}

func testReadFileString(t *testing.T, file string) string {
	t.Helper()
	return string(testReadFile(t, file))
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
