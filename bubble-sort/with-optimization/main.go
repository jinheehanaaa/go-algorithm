package main

import "fmt"

func main() {
	var numbers []int = []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 86, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Our list of numbers is:", numbers)
	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)
}

func bubbleSort(numbers []int) {
	var N int = len(numbers)
	var i int
	for i = 0; i < N; i++ {
		fmt.Println("Doing a sweep:", numbers)
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, prevPasses int) bool {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false

	for secondIndex < (N - prevPasses) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}
