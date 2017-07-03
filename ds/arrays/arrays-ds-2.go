//Arrays-DS
//Problem : https://www.hackerrank.com/challenges/arrays-ds

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var k int
	for i := 1; i <= n; i++ {
		fmt.Scan(&k)
		defer fmt.Printf("%d ", k)
	}
}
