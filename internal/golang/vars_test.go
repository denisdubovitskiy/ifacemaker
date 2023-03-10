package golang

import (
	"go/build"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type envVar struct {
	key string
	val string
}

func TestVars(t *testing.T) {
	cases := []struct {
		name   string
		setenv []*envVar
		fn     func() string
		want   string
	}{
		{
			name: "GOPATH from env",
			setenv: []*envVar{
				{
					key: "GOPATH",
					val: "/path/to/gopath",
				},
			},
			fn:   GOPATH,
			want: "/path/to/gopath",
		},
		{
			name: "GOPATH from build info",
			fn:   GOPATH,
			want: build.Default.GOPATH,
		},
		{
			name: "GOROOT from env",
			setenv: []*envVar{
				{
					key: "GOROOT",
					val: "/path/to/goroot",
				},
			},
			fn:   GOROOT,
			want: "/path/to/goroot",
		},
		{
			name: "GOROOT from build info",
			fn:   GOROOT,
			want: build.Default.GOROOT,
		},
		{
			name: "GOMODCACHE from env",
			setenv: []*envVar{
				{
					key: "GOPATH",
					val: "/path/to/gopath",
				},
			},
			fn:   GOMODCACHE,
			want: "/path/to/gopath/pkg/mod",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_ = os.Unsetenv("GOROOT")     //nolint:errcheck
			_ = os.Unsetenv("GOPATH")     //nolint:errcheck
			_ = os.Unsetenv("GOMODCACHE") //nolint:errcheck

			for _, e := range tc.setenv {
				t.Setenv(e.key, e.val)
			}

			// act
			got := tc.fn()

			// assert
			require.Equal(t, tc.want, got)
		})
	}
}
