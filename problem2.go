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
	conditionOptions := condition(ConditionOptions{MaxResult:4000000})
	fibSequence := fib(conditionOptions)
	sumOfEvenItems := sumOfArrayItems(fibSequence)
	fmt.Println(sumOfEvenItems)
}

func fib(condition func(int, int) bool) []int {
	result := []int{1, 2}
	for i := 2 ; ; i++ {
		currentItem := result[i - 1] + result[i - 2]
		if condition(currentItem, i) == false { break }
		result = append(result, currentItem)
	}
	return result
}

func sumOfArrayItems(items []int) int {
	sum := 0
	for _, value := range items {
		if value % 2 == 0 {
			sum += value
		}
	}
	return sum
}

type ConditionOptions struct {
	MaxResult int
	MaxTerm int
}

func condition(options ConditionOptions) func(int, int) bool {
	return func(currentItem int, currentTerm int) bool {
		if options.MaxTerm == 0 && options.MaxResult != 0 {
			return !(currentItem >= options.MaxResult)
		}
		if options.MaxResult == 0 && options.MaxTerm != 0 {
			return !(currentTerm >= options.MaxTerm)
		}
		return true
	}
}