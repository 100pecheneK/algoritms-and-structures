package main

import (
	"fmt"
	"strings"
)

type Node[T interface{}] struct {
	value T
	next  *Node[T]
}
type Stack[T interface{}] struct {
	node *Node[T]
	Size int
}

func (s *Stack[T]) Push(value T) {
	newNode := new(Node[T])
	newNode.value = value
	if s.node != nil {
		newNode.next = s.node
	}
	s.node = newNode
	s.Size++
}

func (s *Stack[T]) Pop() (res T) {
	if s.node != nil {
		res = s.node.value
		s.node = s.node.next
	}
	s.Size--
	return
}

func (s *Stack[T]) Top() T {
	fmt.Println(s.node.value)
	return s.node.value
}

func main() {

	s1 := "()({[]}())" // true
	s2 := "{}"         // true
	s3 := "}{"         // false
	s4 := ""           // true
	s5 := "(()"        // false
	s6 := "())"        // false
	s7 := "[(])"       // false
	s8 := "}{"

	fmt.Println(isBalanced(s1))
	fmt.Println(isBalanced(s2))
	fmt.Println(isBalanced(s3))
	fmt.Println(isBalanced(s4))
	fmt.Println(isBalanced(s5))
	fmt.Println(isBalanced(s6))
	fmt.Println(isBalanced(s7))
	fmt.Println(isBalanced(s8))
}

func isBalanced(str string) bool {
	var st Stack[rune]

	bks := make(map[rune]rune)
	bks['{'] = '}'
	bks['('] = ')'
	bks['['] = ']'

	for _, s := range str {
		if ok := strings.Contains("[({", string(s)); ok {
			st.Push(s)
		} else if top := st.Pop(); bks[top] != s { // st.size != 0 , st.Pop return default value for rune == ''
			return false
		}
	}
	return st.Size == 0
}
