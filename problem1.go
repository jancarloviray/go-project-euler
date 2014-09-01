/*
If we list all the natural numbers below 10 
that are multiples of 3 or 5, we get 3, 5, 6 and 9. 
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import (
	"fmt"
)

func main(){
	fmt.Println(getSumOfArrayItems(getMultiples(10)))
	fmt.Println(getSumOfArrayItems(getMultiples(1000)))
}

func getMultiples(cap int) []int {
	result := []int{}
	for i := 0 ; i < cap; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			result = append(result, i)
		}
	}
	return result
}

func getSumOfArrayItems(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}