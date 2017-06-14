//Time Conversion
//Problem : https://www.hackerrank.com/challenges/time-conversion

package main

import "fmt"

func main() {

	var h, m, s int
	var str string

	fmt.Scanf("%2d:%2d:%2d%s", &h, &m, &s, &str)

	if str == "PM" && h != 12 {
		h = (h + 12) % 24
	} else if str == "AM" && h == 12 {
		h = 0
	}
	fmt.Printf("%02d:%02d:%02d\n", h, m, s)

}
