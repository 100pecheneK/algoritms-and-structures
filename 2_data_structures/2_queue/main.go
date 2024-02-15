package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}
type Queue[T any] struct {
	node *Node[T]
	size int
}

func (q *Queue[T]) Push(value T) {
	newNode := new(Node[T])
	newNode.value = value
	if q.size != 0 {
		newNode.next = q.node
	}
	q.node = newNode
	q.size++
}

func (q Queue[T]) Pop() (res T) {
	if q.size == 0 {
		return
	}
	if q.size == 1 {
		res = q.node.value
		q.size--
		return
	}

	var newLastNode *Node[T]
	for q.node.next != nil {
		newLastNode = q.node
		q.node = q.node.next
	}
	res = newLastNode.next.value
	newLastNode.next = nil
	q.size--
	return
}

func (q Queue[T]) Front() (res T) {

	for q.node.next != nil {
		q.node = q.node.next
	}
	res = q.node.value
	return
}

func main() {
	q := new(Queue[int])

	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Front())
}
