package parser

type Node struct {
	Token  Token
	Value  any
	Parent *Node
	Left   *Node
	Right  *Node
}
