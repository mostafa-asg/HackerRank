//A Very Big Sum
//Problem : https://www.hackerrank.com/challenges/a-very-big-sum

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	var sum uint64

	for i := 1; i <= n; i++ {
		var tmp uint64
		fmt.Scan(&tmp)
		sum += tmp
	}

	fmt.Println(sum)

}
