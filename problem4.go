/*
A palindromic number reads the same both ways. 
The largest palindrome made from the product of 
two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the 
product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(getPalindromeByDigitCount(2))
	fmt.Println(getPalindromeByDigitCount(3))
}

func getPalindromeByDigitCount(digits int) int {
	largestProduct := 0
	max := ""
	for d := 1 ; d <= digits ; d++ { max += "9" }
	maxInt, _ := strconv.Atoi(max)
	for i := maxInt ; i > 0 ; i-- {
		for j := maxInt ; j > 0 ; j-- {
			product := i * j
			if isPalindrome(product) && product > largestProduct {
				// fmt.Println(product, " : ", i, " x ", j)
				largestProduct = product
			}
		}
	}

	return largestProduct
}

func isPalindrome(num int) bool {
	return num == reverse(num)
}

func reverse(num int) int {
	result := 0
	for num > 0 {
		result = (result * 10) + (num % 10)
		num /= 10
	}
	return result
}