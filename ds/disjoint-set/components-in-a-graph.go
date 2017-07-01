//Components in a graph
//Problem : https://www.hackerrank.com/challenges/components-in-graph

package main

import (
	"fmt"
	"sort"
)

var arr []int
var size []int

func main() {

	var n int
	fmt.Scan(&n)

	setup(n)

	var x, y int

	for i := 1; i <= n; i++ {

		fmt.Scan(&x)
		fmt.Scan(&y)

		pX := parent(x)
		pY := parent(y)

		if pX != pY {
			arr[pY] = pX
			updateSize(pX, pY)
		}
	}

	sort.Ints(size)

	var min int
	for i := 0; i < 2*n+1; i++ {
		if size[i] > 1 {
			min = size[i]
			break
		}
	}

	max := size[len(size)-1]

	fmt.Printf("%d %d\n", min, max)
}

func updateSize(mainSet int, removeSet int) {

	n := size[removeSet]

	size[mainSet] = size[mainSet] + n
	size[removeSet] = 0
}

func parent(index int) int {

	for {
		if arr[index] == index {
			break
		}

		index = arr[index]
	}

	return arr[index]
}

func setup(n int) {

	arr = make([]int, 2*n+1)
	size = make([]int, 2*n+1)

	for i := 0; i < len(arr); i++ {
		arr[i] = i
		size[i] = 1
	}
	size[0] = 0

}
