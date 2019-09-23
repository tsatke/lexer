package ast

type Tree interface {
	Root() Node
}

type Node interface {
	AddChild(Node)
	Evaluate() error
}
