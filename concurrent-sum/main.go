package main

import (
	"fmt"
) 

func sum(numbers []int, ch chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	ch <- sum
}

func main() {
	var length int
	fmt.Println("Enter Numbers want to take as input: ")
	fmt.Scanln(&length)
	num := make([]int, length)
	fmt.Println("Enter Numbers: ")
	for i := range num{
		fmt.Scanln(&num[i])
	}

	ch := make(chan int)

	mid := length/ 2
	left_half := num[:mid]
	right_half := num[mid:]

	go sum(left_half, ch)
	go sum(right_half, ch)

	sum1 := <- ch
	sum2 := <- ch

	total_sum := sum1 + sum2

	fmt.Printf("Total Sum: %d\n", total_sum)

}