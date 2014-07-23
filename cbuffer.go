package main

import (
	"fmt"
)

type CirBuffer struct {
	Maxlen     int
	head, tail int
	overwrite  bool
	q          []int
}

func New(n int) *CirBuffer {
	return &CirBuffer{n, 0, -1, true, make([]int, n)}
}

func (p *CirBuffer) Push(n int) {
	if p.overwrite {
		p.q[p.head%p.Maxlen] = n
	}
	// else {
	// 	// if p.tail%p.Maxlen == p.head%p.Maxlen {

	// 	// 	fmt.Println("Cannot overwrite", n, p.tail, p.head)
	// 	// } else {
	// 	// 	p.q[p.head%p.Maxlen] = n
	// 	// }

	// }

	p.head = p.head + 1
}
func (p *CirBuffer) Pop() (value int) {
	p.tail = (p.tail + 1)
	value = p.q[p.tail%p.Maxlen]

	p.q[p.tail%p.Maxlen] = 0

	return value
}

func main() {

	inputBuffer := New(10)

}
