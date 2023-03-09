package main

import (
	"github.com/Masterminds/semver"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseVersionDirectory(t *testing.T) {
	cases := []struct {
		given string
		want  versionDirectory
	}{
		{
			given: "v5@v5.39.3",
			want: versionDirectory{
				major:   5,
				version: semver.MustParse("v5.39.3"),
			},
		},
		{
			given: "v0.1.0",
			want: versionDirectory{
				major:   0,
				version: semver.MustParse("v0.1.0"),
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.given, func(t *testing.T) {
			t.Parallel()

			parseVersionDirectory(tc.given)
		})
	}
}

func TestSortVersions(t *testing.T) {
	cases := []struct {
		given []string
		want  []string
	}{
		{
			given: []string{
				"v5@v5.39.2",
				"v5@v5.39.3",
			},
			want: []string{
				"v5@v5.39.3",
				"v5@v5.39.2",
			},
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run("", func(t *testing.T) {
			t.Parallel()

			sortVersions(tc.given)

			require.Equal(t, tc.want, tc.given)
		})
	}
}
