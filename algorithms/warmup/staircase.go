//Staircase
//problem : https://www.hackerrank.com/challenges/staircase

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	for row := 1; row <= n; row++ {

		for space := n - row; space > 0; space-- {
			fmt.Print(" ")
		}

		for sharp := 1; sharp <= row; sharp++ {
			fmt.Print("#")
		}

		fmt.Println()

	}

}
