package main

import (
	"container/heap"
	"fmt"
)

type stack struct {
	priority int
	cylinder []int
}

func (s *stack) addCylinder(cyl int) {
	s.cylinder = append(s.cylinder, cyl)
	s.priority += cyl
}

func (s *stack) removeCylinder() {
	s.priority -= s.cylinder[0]
	s.cylinder = s.cylinder[1:]
}

type pqueue []*stack

func (q pqueue) Len() int {
	return len(q)
}

func (q pqueue) Less(i, j int) bool {
	return q[i].priority > q[j].priority
}

func (q pqueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *pqueue) Push(x interface{}) {
	*q = append(*q, x.(*stack))
}

func (q *pqueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func isEqual(stacks pqueue) (int, bool) {

	for _, v := range stacks {
		if v.priority == 0 {
			return 0, true
		}
	}

	for i := 1; i < len(stacks); i++ {
		if stacks[i].priority != stacks[i-1].priority {
			return 0, false
		}
	}

	return stacks[0].priority, true

}

func main() {

	var n1, n2, n3 int
	fmt.Scan(&n1, &n2, &n3)

	var cyl int
	var q pqueue

	heap.Init(&q)

	s := new(stack)
	for i := 1; i <= n1; i++ {
		fmt.Scan(&cyl)
		s.addCylinder(cyl)
	}
	heap.Push(&q, s)

	s = new(stack)
	for i := 1; i <= n2; i++ {
		fmt.Scan(&cyl)
		s.addCylinder(cyl)
	}
	heap.Push(&q, s)

	s = new(stack)
	for i := 1; i <= n3; i++ {
		fmt.Scan(&cyl)
		s.addCylinder(cyl)
	}
	heap.Push(&q, s)

	for {
		answer, equal := isEqual(q)
		if equal {
			fmt.Println(answer)
			break
		}

		s = heap.Pop(&q).(*stack)
		s.removeCylinder()
		heap.Push(&q, s)
	}
}
