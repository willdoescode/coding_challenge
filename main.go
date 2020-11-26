package main

import "fmt"

type Node struct {
	elem int
	next *Node
}

func (n *Node) display() []int {
	res := make([]int, 0)
	res = append(res, n.elem)
	if n.next != nil {
		res = append(res, n.next.display()...)
	}

	return res
}

func (n *Node) append(elem int) {
	if n.next == nil {
		n.next = &Node{elem: elem, next: nil}
		return
	}
	n.next.append(elem)
}

func append_front(n Node, elem int) Node {
	return Node{elem: elem, next: &n}
}

func (n *Node) find(num int, state int) int {
	if n.elem == num {
		return state
	}
	state = state + 1
	return n.next.find(num, state)
}

func main() {
	h := Node{elem: 1, next: nil}
	for i := 2; i < 20; i++ {
		h.append(i)
	}
	fmt.Println(h.display())
	fmt.Println(h.find(5, 0))
}
