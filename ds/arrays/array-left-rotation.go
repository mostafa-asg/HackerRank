//Left Rotation
//Problem : https://www.hackerrank.com/challenges/array-left-rotation

package main

import (
	"fmt"
)

func main() {

	var n, d int

	fmt.Scan(&n)
	fmt.Scan(&d)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[(i+n-d)%n])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}

}
