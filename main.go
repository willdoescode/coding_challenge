package main

import "fmt"

func main() {
	h := Node{elem: 4, next: nil}
	h.append(6)
	h.append(3)
	h.append_front(5)
	fmt.Println(h.display())
	//h.append_range(1, 20)
	//fmt.Println(h.display())
	//h.rm_index(5)
	//fmt.Println(h.display())
	//h.rm_num(3)
	//fmt.Println(h.display())
	//h.append_slice([]int{1, 2, 4, 3})
	//fmt.Println(h.display())
	//also := list_from_slice([]int{1, 3, 2, 5})
	//fmt.Println(also.display())
}
