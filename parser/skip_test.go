package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Skip(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		want  *Node
		error error
	}{
		{
			name: "empty",
			in:   "",
			want: nil,
		},
		{
			name:  "invalid literal",
			in:    "$skip=3.5",
			error: ParseErr,
		},
		{
			name:  "missing literal",
			in:    "$skip=",
			error: ParseErr,
		},
		{
			name: "valid literal",
			in:   "$skip=3",
			want: &Node{
				Token: Int,
				Value: 3,
			},
		},
		{
			name: "missing $",
			in:   "skip=3",
			want: &Node{
				Token: Int,
				Value: 3,
			},
		},
		{
			name: "upper and lower casing",
			in:   "$sKiP=3",
			want: &Node{
				Token: Int,
				Value: 3,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			node, err := ParsePlaceholder(test.in)
			assert.ErrorIs(t, err, test.error)
			assert.Equal(t, test.want, node)
		})
	}
}
