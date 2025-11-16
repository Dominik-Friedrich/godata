package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Top(t *testing.T) {
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
			in:    "$top=3.5",
			error: ParseErr,
		},
		{
			name:  "missing literal",
			in:    "$top=",
			error: ParseErr,
		},
		{
			name: "valid literal",
			in:   "$top=3",
			want: &Node{
				Token: Int,
				Value: 3,
			},
		},
		{
			name: "missing $",
			in:   "top=3",
			want: &Node{
				Token: Int,
				Value: 3,
			},
		},
		{
			name: "upper and lower casing",
			in:   "$ToP=3",
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
