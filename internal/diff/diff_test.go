package diff

import (
	"slices"
	"testing"
)

func TestDiff(t *testing.T) {
	tests := []struct {
		name string
		a, b []string
		want []Edit
	}{
		{
			name: "identical",
			a:    []string{"x", "y"},
			b:    []string{"x", "y"},
			want: []Edit{{Equal, "x"}, {Equal, "y"}},
		},
		{
			name: "both empty",
			a:    nil,
			b:    nil,
			want: nil,
		},
		{
			name: "old empty",
			a:    nil,
			b:    []string{"x"},
			want: []Edit{{Insert, "x"}},
		},
		{
			name: "new empty",
			a:    []string{"x"},
			b:    nil,
			want: []Edit{{Delete, "x"}},
		},
		{
			name: "changed line",
			a:    []string{"x", "y"},
			b:    []string{"x", "z"},
			want: []Edit{{Equal, "x"}, {Delete, "y"}, {Insert, "z"}},
		},
		{
			name: "line changed to blank",
			a:    []string{"x"},
			b:    []string{""},
			want: []Edit{{Delete, "x"}, {Insert, ""}},
		},
		{
			name: "blank line appended",
			a:    []string{"x"},
			b:    []string{"x", ""},
			want: []Edit{{Equal, "x"}, {Insert, ""}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diff(tt.a, tt.b)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Diff(%q, %q)\n got: %v\nwant: %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
