//simple-text-editor
//Problem : https://www.hackerrank.com/challenges/simple-text-editor

package main

import "fmt"

type Stack []string

func main() {

	var n int
	fmt.Scan(&n)

	editor := ""
	var stack Stack
	var op int
	var str string

	for i := 1; i <= n; i++ {

		fmt.Scan(&op)

		switch op {
		case 1:
			fmt.Scan(&str)
			stack = append(stack, editor)
			editor += str
		case 2:
			fmt.Scan(&op)
			stack = append(stack, editor)
			editor = editor[0 : len(editor)-op]
		case 3:
			fmt.Scan(&op)
			fmt.Println(string(editor[op-1]))
		case 4:
			size := len(stack)
			if size > 0 {
				editor = stack[size-1]
				stack = stack[:size-1]
			}
		}

	}
}
