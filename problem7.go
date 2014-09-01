/*
By listing the first six prime numbers: 
2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package main 

import "fmt"

func main(){
	// TODO: currently a brute force method with inefficient algorithm. Improve
	fmt.Println(getNthPrimeNumber(10001)) // 104743
}

func getNthPrimeNumber(cap int) int {
	bucket := []int{}

	for i := 2 ; len(bucket) < cap ; i++ {
		if isPrimeNumber(i) {
			bucket = append(bucket, i)
		}
	}

	fmt.Println(bucket)

	return bucket[len(bucket) - 1]
}

func isPrimeNumber(num int) bool {
	for i := 2 ; i < num ; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}
