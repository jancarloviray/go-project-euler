/*
Each new term in the Fibonacci sequence is generated 
by adding the previous two terms. By starting with 1 and 2, 
the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence 
whose values do not exceed four million, 
find the sum of the even-valued terms.
*/

package main

import "fmt"

func main(){
	fmt.Println(fib(10))
}

func fib(terms int) []int {
	result := []int{1, 2}
	for i := 2; i < terms; i++ {
		result = append(result, result[i - 1] + result[i - 2])
	}
	return result
}