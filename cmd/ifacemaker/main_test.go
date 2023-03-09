package main

import (
	"fmt"
	"github.com/Masterminds/semver"
	"testing"

	"github.com/stretchr/testify/require"
)

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

			sortVersions(tc.given)

			require.Equal(t, tc.want, tc.given)
		})
	}
}

func TestParseModule(t *testing.T) {
	cases := []struct {
		name    string
		version string
		want    *sourcePackage
	}{
		{
			name:    "github.com/mattermost/mattermost-server/v5",
			version: "",
			want: &sourcePackage{
				Name: "github.com/mattermost/mattermost-server",
				Base: "mattermost-server",
				Dir:  "github.com/mattermost/mattermost-server",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mod, err := parseModule(tc.name, "")

			require.NoError(t, err)
			fmt.Println(mod)
		})
	}
}
