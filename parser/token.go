package parser

import "unicode"

// Token is the set of lexical tokens of the OData protocol
// Heavily based on https://cs.opensource.google/go/go/+/refs/tags/go1.25.4:src/go/token/token.go
type Token int

const (
	// Special tokens
	Illegal Token = iota

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	Identifier // main
	Int        // 12345
	Float      // 123.45
	String     // 'abc'
	Duration   // duration'P12DT23H59M59.999999999999S'
	TimeOfDay  // 07:59:59.999
	literal_end

	operator_beg
	// Comparison operators
	Equal          // eq
	NotEqual       // ne
	Greater        // gt
	GreaterOrEqual // ge
	Less           // lt
	LessOrEqual    // le
	Has            // has
	In             // in

	// Logical operators
	And // and
	Or  // or
	Not // not

	// Arithmetic Operators
	Add        // add
	Sub        // sub
	Mul        // mul
	Div        // div
	DecimalDiv // divby
	Mod        // mod

	// Grouping Operators
	LParen // (
	RParen // )

	// Other
	LBrack // [
	LBrace // {
	Comma  // ,
	Period // .

	RBrack    // ]
	RBrace    // }
	Semicolon // ;
	Colon     // :

	operator_end

	function_beg
	// Built-In Functions

	// String and Collection Functions
	Concat     // concat(concat(City,', '), Country) eq 'Berlin, Germany'
	Contains   // contains(CompanyName,'freds')
	EndsWith   // endswith(CompanyName,'Futterkiste')
	IndexOf    // indexof(CompanyName,'lfreds') eq 1
	Length     // length(CompanyName) eq 19
	StartsWith // startswith(CompanyName,’Alfr’)
	Substring  // substring(CompanyName,1) eq 'lfreds Futterkiste'

	// Collection Functions
	HasSubset      // hassubset([4,1,3],[3,1])
	HasSubsequence // hassubsequence([4,1,3,1],[1,1])

	// String Functions
	MatchesPattern // matchesPattern(CompanyName,'%5EA.*e$')
	ToLower        // tolower(CompanyName) eq 'alfreds futterkiste'
	ToUpper        // toupper(CompanyName) eq 'ALFREDS FUTTERKISTE'
	Trim           // trim(CompanyName) eq 'Alfreds Futterkiste'

	// Type Functions
	Cast //	cast(ShipCountry,Edm.String)
	IsOf // isof(NorthwindModel.Order)
	// FIXME: Why does the doc define this twice?
	// IsOf // isof(ShipCountry,Edm.String)

	// Geo Functions
	GeoDistance   // geo.distance(CurrentPosition,TargetPosition)
	GeoIntersects // geo.intersects(Position,TargetArea)
	GeoLength     // geo.length(DirectRoute)

	// Conditional Functions
	Case // case(X gt 0:1,X lt 0:-1,true:0)
	function_end

	keyword_beg
	// Keywords
	Asc
	Desc
	keyword_end
)

var keywords map[string]Token

func init() {
	keywords = map[string]Token{
		"asc":  Asc,
		"desc": Desc,
	}
}

// Lookup maps an identifier to its keyword token or [IDENT] (if not a keyword).
func Lookup(ident string) Token {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return Identifier
}

// Predicates

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
func (tok Token) IsLiteral() bool { return literal_beg < tok && tok < literal_end }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
func (tok Token) IsOperator() bool {
	return operator_beg < tok && tok < operator_end
}

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
func (tok Token) IsKeyword() bool { return keyword_beg < tok && tok < keyword_end }

// IsFunction returns true for tokens corresponding to built-in functions; it returns false otherwise.
func (tok Token) IsFunction() bool {
	return function_beg < tok && tok < function_end
}

// IsKeyword reports whether name is a OData keyword, such as "asc" or "desc".
func IsKeyword(name string) bool {
	_, ok := keywords[name]
	return ok
}

// IsIdentifier reports whether name is an identifier, that is, a non-empty
// string made up of letters, digits, and underscores, where the first character
// is not a digit. Keywords are not identifiers.
func IsIdentifier(name string) bool {
	if name == "" || IsKeyword(name) {
		return false
	}
	for i, c := range name {
		if !unicode.IsLetter(c) && c != '_' && (i == 0 || !unicode.IsDigit(c)) {
			return false
		}
	}
	return true
}
