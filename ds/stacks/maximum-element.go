//Maximum Element
//Problem : https://www.hackerrank.com/challenges/maximum-element

package main

import "fmt"

type StackItem struct {
	val int
	max int
}

func main() {

	var stack []StackItem
	max := 0
	var data int
	var n int

	fmt.Scan(&n)

	for n > 0 {

		fmt.Scan(&data)

		switch data {

		case 1:
			fmt.Scan(&data)

			if data > max {
				max = data
			}

			item := StackItem{data, max}

			stack = append(stack, item)
		case 2:
			size := len(stack)
			if size > 0 {
				stack = stack[0 : size-1]
			}

			size = len(stack)
			if size > 0 {
				max = stack[size-1].max
			} else {
				max = 0
			}
		case 3:
			size := len(stack)
			if size > 0 {
				fmt.Println(stack[size-1].max)
			}
		}

		n--

	}

}
