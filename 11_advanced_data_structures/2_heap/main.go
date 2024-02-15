package main

import (
	"fmt"
)

type Heap []int

// x>>1 = x/=2
func (h *Heap) add(x int) { // O(log(n))
	*h = append(*h, x)
	ind := len(*h) - 1
	for ind != 1 && (*h)[ind] < (*h)[ind>>1] {
		(*h)[ind], (*h)[(ind>>1)] = (*h)[ind>>1], (*h)[ind]
		ind >>= 1
	}
}
func (h *Heap) erase() { // O(log(n))
	(*h)[1], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[1]
	*h = (*h)[:len(*h)-1]
	ind := 1
	for (ind<<1) < len(*h) && (*h)[ind] > (*h)[(ind<<1)] || (ind<<1)+1 < len(*h) && (*h)[ind] > (*h)[(ind<<1)+1] {
		if (ind<<1)+1 >= len(*h) || (*h)[(ind<<1)] < (*h)[(ind<<1)+1] {
			(*h)[ind], (*h)[(ind<<1)] = (*h)[(ind<<1)], (*h)[ind]
			ind <<= 1
		} else {
			(*h)[ind], (*h)[(ind<<1)+1] = (*h)[(ind<<1)+1], (*h)[ind]
			ind <<= 1
			ind++
		}
	}
}
func (h Heap) top() int {
	return h[1]
}

func main() {
	heap := &Heap{0, 1}
	heap.add(2)
	heap.add(5)
	heap.add(3)
	heap.add(10)
	heap.add(6)
	heap.add(9)
	heap.add(5)
	heap.add(12)
	heap.add(40)

	fmt.Println(heap)
	fmt.Println(heap.top())
	heap.erase()
	fmt.Println(heap.top())
	heap.erase()
	fmt.Println(heap.top())
	heap.erase()
	fmt.Println(heap.top())
	heap.erase()
	fmt.Println(heap.top())
	heap.erase()
	fmt.Println(heap.top())
	heap.erase()
	fmt.Println(heap.top())
	heap.erase()
	fmt.Println(heap.top())
	heap.erase()
	fmt.Println(heap.top())
	heap.erase()
	fmt.Println(heap.top())

}
