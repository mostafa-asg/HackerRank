// problem :
// https://www.hackerrank.com/challenges/simple-array-sum

package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanln(&n)

	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	fmt.Println(sum)

}
