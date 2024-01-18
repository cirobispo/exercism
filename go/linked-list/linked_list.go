package linkedlist

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

type Node struct {
	Value interface{}
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
	size int
} 

func NewList(elements ...interface{}) *List {
	result:=&List{}
	for i:=range elements {
		result.Push(elements[i])
	}

	return result
}

func (n *Node) Next() *Node {
	if n != nil {
		return n.next
	}
	return nil
}

func (n *Node) Prev() *Node {
	if n != nil {
		return n.prev
	}
	return nil
}

func newNode(prev *Node, value interface{}) *Node {
	return &Node{prev: prev, Value: value}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(v interface{}) {
	node:=newNode(l.tail, v)
	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node

	if l.head == nil {
		l.head = node
	}
	l.size++
}

func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		return nil, nil
	}

	value := l.tail.Value
	if l.tail.prev != nil {
		l.tail=l.tail.prev
		l.tail.next=nil
	} else {
		l.tail=nil
		l.head=nil
	}
	l.size--
	return value, nil
}

func (l *List) Unshift(v interface{}) {
	node:=newNode(nil, v)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}

	l.size++
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, nil
	}
	value:=l.head.Value
	if l.head.next != nil {
		l.head=l.head.next
		l.head.prev=nil
	} else {
		l.head=nil
		l.tail=nil
	}
	l.size--

	return value, nil
}

func (l *List) Reverse() {
	if l.head != nil {
		node:=l.head.Next()
		for node != nil {
			temp:=node
			node = node.Next()
			temp.next, temp.prev=temp.prev, temp.next
		}
		temp:=l.head
		l.head=l.tail
		temp.next,temp.prev=temp.prev,temp.next
		l.tail=temp
	}
}
