package linkedlist

import "fmt"

// Define the List and Element types here.
type node struct {
	value    int
	neighbor *node
}

type List struct {
	tail *node
	size int
}

func New(elements []int) *List {
	result := &List{}
	for i := range elements {
		result.Push(elements[i])
	}

	return result
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	next := &node{value: element}
	if l.tail != nil {
		next.neighbor = l.tail
	}

	l.tail = next
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.tail == nil {
		return -1, fmt.Errorf("list is empty")
	}

	result := l.tail.value
	l.tail = l.tail.neighbor
	l.size--

	return result, nil
}

func (l *List) Array() []int {
	result := make([]int, l.size)

	idx := l.size - 1
	last := l.tail
	for last != nil {
		result[idx] = last.value
		last = last.neighbor
		idx--
	}

	return result
}

func (l *List) Reverse() *List {
	if l.tail == nil {
		return l
	}

	var child *node
	for {
		parent := l.tail.neighbor
		l.tail.neighbor = child
		child = l.tail
		if parent == nil {
			break
		}
		l.tail = parent
	}

	return l
}
