//Plus Minus
//https://www.hackerrank.com/challenges/plus-minus

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var (
		positives int
		negatives int
		zeros     int
	)

	for _, num := range arr {

		if num > 0 {
			positives++
		} else if num == 0 {
			zeros++
		} else {
			negatives++
		}

	}

	fmt.Printf("%0.6f\n", float32(positives)/float32(n))
	fmt.Printf("%0.6f\n", float32(negatives)/float32(n))
	fmt.Printf("%0.6f\n", float32(zeros)/float32(n))
}
