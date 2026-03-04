package utils_test

import (
	"github.com/universtar-org/ust/internal/utils"
	"testing"
)

func TestParseOwner(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path string
		want string
	}{
		{
			name: "NormalOwner",
			path: "alice.yaml",
			want: "alice",
		},
		{
			name: "UppercaseOwner",
			path: "BOB.YAML",
			want: "BOB",
		},
		{
			name: "OwnerWithDash",
			path: "foo-bar.yaml",
			want: "foo-bar",
		},
		{
			name: "OwnerWithDots",
			path: "foo.foo.bar.bar.yaml",
			want: "foo.foo.bar.bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.ParseOwner(tt.path)
			utils.HandleTestDiff(t, tt.want, got)
		})
	}
}
