package mdtopdf

import (
	bf "gopkg.in/russross/blackfriday.v2"
)

type listType int

const (
	notlist listType = iota
	unordered
	ordered
	definition
)

func (n listType) String() string {
	switch n {
	case notlist:
		return "Not a List"
	case unordered:
		return "Unordered"
	case ordered:
		return "Ordered"
	case definition:
		return "Definition"
	}
	return ""
}

type containerState struct {
	containerType  bf.NodeType
	textStyle      Styler
	leftMargin     float64
	firstParagraph bool

	// populated if node type is a list
	listkind   listType
	itemNumber int // only if an ordered list

	// populated if node type is a link
	destination string
}

type states struct {
	stack []*containerState
}

func (s *states) push(c *containerState) {
	s.stack = append(s.stack, c)
}

func (s *states) pop() *containerState {
	var x *containerState
	x, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
	return x
}

func (s *states) peek() *containerState {
	return s.stack[len(s.stack)-1]
}

func (s *states) parent() *containerState {
	return s.stack[len(s.stack)-2]
}
