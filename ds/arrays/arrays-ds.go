//Arrays-DS
//Problem : https://www.hackerrank.com/challenges/arrays-ds

package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
}
