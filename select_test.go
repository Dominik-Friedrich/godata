package godata

import (
	"fmt"
	"godata/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SelectParsePlaceholder(in string) (*Select, error) {
	return nil, fmt.Errorf("%w: unimplemented", parser.ParseErr)
}

func TestParse_Select(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		want  *Select
		error error
	}{
		{
			name: "empty",
			in:   "",
			want: nil,
		},
		{
			name: "upper and lower casing",
			in:   "$sEleCt=Rating",
			want: &Select{
				Items: []SelectItem{
					{
						Property: &SelectProperty{
							Name: "Rating",
						},
					},
				},
			},
		},
		{
			name: "5.1.3 Select - simple",
			in:   "$select=Rating,ReleaseDate",
			want: &Select{
				Items: []SelectItem{
					{
						Property: &SelectProperty{
							Name: "Rating",
						},
					},
					{
						Property: &SelectProperty{
							Name: "ReleaseDate",
						},
					},
				},
			},
		},
		{
			name: "5.1.3 Select - $ is optional",
			in:   "select=Rating,ReleaseDate",
			want: &Select{
				Items: []SelectItem{
					{
						Property: &SelectProperty{
							Name: "Rating",
						},
					},
					{
						Property: &SelectProperty{
							Name: "ReleaseDate",
						},
					},
				},
			},
		},
		{
			name: "5.1.3 Select - with star",
			in:   "select=*",
			want: &Select{
				Items: []SelectItem{
					{
						SelectAll: true,
					},
				},
			},
		},
		{
			name: "5.1.3 Select - with property of complex property",
			in:   "select=Address/Country",
			want: &Select{
				Items: []SelectItem{
					{
						Property: &SelectProperty{
							Name: "Address",
							Sub: &SelectProperty{
								Name: "Country",
							},
						},
					},
				},
			},
		},
		{
			name: "5.1.3 Select - nested select",
			// Original, adjusted for reduced feature set
			// $select=Address($select=Street,City,Namespace.AddressWithLocation/Location)
			in: "$select=Address($select=Street,City)",
			want: &Select{
				Items: []SelectItem{
					{
						Property: &SelectProperty{
							Name: "Address",
							Options: &SelectOptions{
								Select: &Select{
									Items: []SelectItem{
										{
											Property: &SelectProperty{
												Name: "Street",
											},
										},
										{
											Property: &SelectProperty{
												Name: "City",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "5.1.3 Select - nested filter etc.",
			// Original, adjusted for reduced feature set
			// $select=Addresses($filter=startswith(City,'H');$top=5;$skip=0;$count=true;$orderby=$it;$search=blue;@c=15)&$expand=Addresses/Country
			in: "$select=Addresses($filter=startswith(City,'H');$top=5;$skip=0;$count=true;$orderby=City)",
			want: &Select{
				Items: []SelectItem{
					{
						Property: &SelectProperty{
							Name: "Addresses",
							Options: &SelectOptions{
								Filter: &Filter{
									// TODO implemente filter
								},
								Top:   Ptr[Top](5),
								Skip:  Ptr[Skip](0),
								Count: Ptr[Count](true),
								OrderBy: &OrderBy{
									Items: []OrderByItem{
										{
											Property: "City",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sel, err := SelectParsePlaceholder(test.in)

			if test.error != nil {
				assert.ErrorIs(t, err, test.error)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want, sel)
		})
	}
}
