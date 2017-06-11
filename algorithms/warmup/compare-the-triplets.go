package main

import "fmt"

func main() {

	alice := 0
	bob := 0

	aliceTriple := make([]int, 3)
	bobTriple := make([]int, 3)

	for i := 0; i < 3; i++ {
		fmt.Scan(&aliceTriple[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Scan(&bobTriple[i])
	}

	for i := 0; i < 3; i++ {
		compare(aliceTriple[i], bobTriple[i], &alice, &bob, increasePoint)
	}

	fmt.Printf("%d %d\n", alice, bob)

}

func compare(v1 int, v2 int, alice *int, bob *int, givePointTo func(p *int)) {

	if v1 > v2 {
		givePointTo(alice)
	}

	if v2 > v1 {
		givePointTo(bob)
	}

}

func increasePoint(point *int) {

	*point = *point + 1

}
