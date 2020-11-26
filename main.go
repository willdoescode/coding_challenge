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

func (n *Node) find(num int, offset int) int {
	offset++
	if n.elem == num {
		return offset
	}
	offset++
	n.next.find(num, offset)
	return offset
}

func main() {
	h := Node{elem: 1, next: nil}
	fmt.Println(h.display())
	h.append(5)
	fmt.Println(h.display())
	h = append_front(h, 6)
	fmt.Println(h.display())
	fmt.Println(h.find(5, 0))
}
