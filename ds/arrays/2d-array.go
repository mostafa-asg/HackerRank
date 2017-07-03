//2D Array - DS
// Problem : https://www.hackerrank.com/challenges/2d-array
package main

import (
	"fmt"
	"sync"
)

const size = 6
var producer sync.WaitGroup
var consumer sync.WaitGroup
var max int

func main() {

	max = 7*-9
	var arr [size][size]int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	ch := make(chan int, 16)
	consumer.Add(1)

	for i:=0;i<=size/2;i++ {
		for j:=0;j<=size/2;j++ {
			producer.Add(1)
			go calulate(arr,i,j,ch)
		}
	}

	go findMax( ch )

	producer.Wait()
	close(ch)

	consumer.Wait()
	fmt.Println( max )
}

func findMax(ch <-chan int) {

	for {

		val , more := <-ch

		if( more ) {
			if val>max {
				max = val
			}
		} else {
			consumer.Done()
			break
		}

	}

}

func calulate(arr [size][size]int , row int , col int , ch chan<- int) {

	defer producer.Done()

	sum := 0

	for i:=row;i<row+size/2;i++ {
		for j:=col;j<col+size/2;j++ {
			sum += arr[i][j]
		}
	}

	sum -= arr[row+1][col]
	sum -= arr[row+1][col+2]

	ch <- sum
}
