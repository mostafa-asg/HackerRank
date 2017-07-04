// Sparse arrays
// Prblem : https://www.hackerrank.com/challenges/sparse-arrays
package main

import (
	"fmt"
)

func main() {

	var n int
	var str string
	wordMap := make(map[string]int)

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		fmt.Scan(&str)

		addToMap(wordMap, str)

	}

	var q int
	fmt.Scan(&q)

	for i := 1; i <= q; i++ {
		fmt.Scan(&str)

		fmt.Println(getFromMap(wordMap, str))
	}

}

func getFromMap(wordMap map[string]int, str string) int {

	count, exists := wordMap[str]

	if exists {
		return count
	} else {
		return 0
	}

}

func addToMap(wordMap map[string]int, str string) {

	count, exists := wordMap[str]

	if exists {
		wordMap[str] = count + 1
	} else {
		wordMap[str] = 1
	}

}
