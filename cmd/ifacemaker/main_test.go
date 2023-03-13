package main

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestFindSourceFiles(t *testing.T) {
	t.Parallel()

	finder := newSourceFilesFinder()
	finder.fs = afero.NewMemMapFs()

	files1 := []string{
		"directory_1/README.md",
		"directory_1/go.mod",
		"directory_1/go.sum",
		"directory_1/source_code.go",
		"directory_1/source_code_test.go",
		"directory_1/subdir/source_code.go",
		"directory_1/subdir/source_code_test.go",
	}
	files2 := []string{
		"directory_2/README.md",
		"directory_2/go.mod",
		"directory_2/go.sum",
		"directory_2/source_code.go",
		"directory_2/source_code_test.go",
		"directory_2/subdir/source_code.go",
		"directory_2/subdir/source_code_test.go",
	}
	files := append(files1, files2...)

	for _, f := range files {
		_ = afero.WriteFile(finder.fs, f, []byte(""), os.ModePerm) //nolint:errcheck
	}

	cases := []struct {
		directory string
		want      []string
	}{
		{"directory_1", []string{"directory_1/source_code.go"}},
		{"directory_2", []string{"directory_2/source_code.go"}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.directory, func(t *testing.T) {
			t.Parallel()

			// act
			got, err := finder.findSourceFiles("directory_1")

			// assert
			require.NoError(t, err)
			require.Equal(t, []string{"directory_1/source_code.go"}, got)
		})
	}
}
