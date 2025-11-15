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
			name:  "empty",
			in:    "",
			error: parseErr, // TODO: to err or not to err
		},
		{
			name:  "invalid literial",
			in:    "$skip=3.5",
			error: parseErr,
		},
		{
			name:  "no literial",
			in:    "$skip=",
			error: parseErr,
		},
		{
			name: "valid literial",
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
			node, err := parsePlaceholder(test.in)
			assert.ErrorIs(t, err, test.error)
			assert.Equal(t, test.want, node)
		})
	}
}
