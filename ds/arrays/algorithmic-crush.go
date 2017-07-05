//Algorithmic Crush
//Problem : https://www.hackerrank.com/challenges/crush
package main

import (
	"fmt"
)

func main() {

	var n, m, a, b, k, i int32
	var max int64 = 0

	fmt.Scan(&n, &m)

	arr := make([]int64, n+1)

	for i = 1; i <= m; i++ {

		fmt.Scan(&a, &b, &k)

		arr[a] += int64(k)
		if b+1 <= n {
			arr[b+1] -= int64(k)
		}
	}

	var sum int64
	for i = 1; i <= n; i++ {
		sum += arr[i]
		if sum > max {
			max = sum
		}
	}

	fmt.Println(max)

}
