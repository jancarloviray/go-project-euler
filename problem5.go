/*
2520 is the smallest number that can be divided 
by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is 
evenly divisible by all of the numbers from 1 to 20?
*/

package main

import "fmt"

func main(){
	fmt.Println(getSmallestNumberByCondition(isDivisibleBy1to20))
}

type condition func(int) bool

func getSmallestNumberByCondition(c condition) int {
	result := 0
	for i := 20 ; result == 0 ; i++ {
		if c(i) { result = i }
	}
	return result
}

func isDivisibleBy1to20(num int) bool {
	for i := 1 ; i <= 20 ; i++ {
		if num % i != 0 { return false }
	}
	return true
}