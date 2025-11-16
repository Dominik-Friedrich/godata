package godata

import "godata/parser"

type ServiceRequest[T QueryOptions] struct {
	Service string
	Request T
}

type QueryOptions struct {
	// Schemaversion Schemaversion
	Apply *Apply
	//Compute *Compute
	Filter  *Filter
	Count   *Count
	OrderBy *OrderBy
	Skip    *Skip
	Top     *Top
	Expand  *Expand
	//Search *Search
	Select *Select
	//Format *Format
}

type Apply struct{}

type Filter struct {
	Items []parser.Node
}

// The Count system query option with a value of true specifies that the total count
// of items within a collection matching the request be returned along with the result.
//
// The $count system query option ignores any Top, Skip, or Expand query options,
// and returns the total count of results across all pages including only those results matching any specified
// Filter and Search.
//
// https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html#sec_SystemQueryOptioncount
type Count bool

// OrderBy represents an $orderby query parameter
//
// Partial definition:
// orderby     = ( "$orderby" / "orderby" ) EQ orderbyItem *( COMMA orderbyItem )
// orderbyItem = commonExpr [ RWS ( "asc" / "desc" ) ]
//
// https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html#sec_SystemQueryOptionorderby
type OrderBy struct {
	Items []OrderByItem
}

type OrderByItem struct {
	Desc     bool // True: desc, False(default): asc
	Property string
}

// The Skip system query option specifies a non-negative integer n that excludes the first n items
// of the queried collection from the result. The service returns items starting at position n+1.
//
// https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html#sec_SystemQueryOptionskip
type Skip uint

// The Top system query option specifies a non-negative integer n that limits the number of items returned from a collection.
// The service returns the number of available items up to but not greater than the specified value n.
//
// https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html#sec_SystemQueryOptiontop
type Top uint

type Expand struct{}

type Select struct {
	Items []SelectItem
}

// SelectItem represents a single SelectItem in the comma separated Select list.
type SelectItem struct {
	SelectAll    bool
	Property     *SelectProperty
	QueryOptions *QueryOptions
}

type SelectProperty struct {
	Name    string
	Options *SelectOptions
	Sub     *SelectProperty // for "/" selectProperty
}

type SelectOptions struct {
	//Compute *Compute
	Filter *Filter
	// Search *Search
	OrderBy *OrderBy
	Skip    *Skip
	Top     *Top
	Select  *Select
	Count   *Count
}
