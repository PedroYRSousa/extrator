package modules

import (
	"os"
	"testing"
)

func TestGetDataPaths(t *testing.T) {
	want := os.TempDir()
	got := GetDataPaths()
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestJoinPaths(t *testing.T) {
	tests := []struct {
		name     string
		segments []string
		want     string
	}{
		{
			name:     "no segments",
			segments: []string{},
			want:     ".",
		},
		{
			name:     "single relative",
			segments: []string{"foo"},
			want:     "foo",
		},
		{
			name:     "multiple relative",
			segments: []string{"foo", "bar", "baz"},
			want:     "foo/bar/baz",
		},
		{
			name:     "absolute path",
			segments: []string{"/foo", "bar"},
			want:     "/foo/bar",
		},
		{
			name:     "dot segments",
			segments: []string{"foo", ".", "bar"},
			want:     "foo/bar",
		},
		{
			name:     "dotdot removes previous",
			segments: []string{"foo", "bar", "..", "baz"},
			want:     "foo/baz",
		},
		{
			name:     "dotdot at start relative",
			segments: []string{"..", "foo"},
			want:     "../foo",
		},
		{
			name:     "dotdot at start absolute",
			segments: []string{"/", "..", "foo"},
			want:     "/foo",
		},
		{
			name:     "empty segments ignored",
			segments: []string{"", "foo", "", "bar"},
			want:     "foo/bar",
		},
		{
			name:     "preserve trailing slash",
			segments: []string{"foo/"},
			want:     "foo/",
		},
		{
			name:     "preserve trailing slash absolute",
			segments: []string{"/foo/"},
			want:     "/foo/",
		},
		{
			name:     "absolute resets path",
			segments: []string{"foo", "/bar", "baz"},
			want:     "/bar/baz",
		},
		{
			name:     "result empty absolute becomes root",
			segments: []string{"/"},
			want:     "/",
		},
		{
			name:     "result empty relative becomes dot",
			segments: []string{"."},
			want:     ".",
		},
		{
			name:     "only double dot relative",
			segments: []string{".."},
			want:     "..",
		},
		{
			name:     "multiple double dots relative",
			segments: []string{"..", "..", "foo"},
			want:     "../../foo",
		},
		{
			name:     "reset absolute in middle",
			segments: []string{"foo", "bar", "/baz", "qux"},
			want:     "/baz/qux",
		},
		{
			name:     "trailing slash lost if not last non-empty",
			segments: []string{"foo/", "bar"},
			want:     "foo/bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinPaths(tt.segments...)
			if got != tt.want {
				t.Fatalf("segments %v: expected %q, got %q", tt.segments, tt.want, got)
			}
		})
	}
}
