/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
	"math"
)

func main(){
	// fmt.Println(getBiggestPrime(13195))
	fmt.Println(getBiggestPrime(600851475143))
}

func getBiggestPrime(num int) int {
	max := math.Sqrt(float64(num))
	for i := int(max) ; i >= 2 ; i-- {
		fmt.Println(i)
		if num % i == 0 && getBiggestPrime(i) == 1 {
			return i
		}
	}
	return 1
}