//Birthday Cake Candles
//Problem : https://www.hackerrank.com/challenges/birthday-cake-candles

package main

import (
	"fmt"
	"sort"
)

func main() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	count := 0

	for i := 0; i < n; i++ {
		count++
		if i != len(arr)-1 && arr[i] != arr[i+1] {
			break
		}
	}

	fmt.Println(count)

}
