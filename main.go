package main

import (
	"fmt"
	"math"
)

// NICHT FÜR MULTITHREADING GEEIGNET :(

func main() {
	fmt.Printf("%d", findNumberByBouncyPercentage(0.99))

}

func returnDigit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func countDigits(num int) (count int) {
	for num != 0 {
		num /= 10
		count = count + 1
	}
	return count
}

func isBouncy(num int) bool {
	var isIncreasing = true
	var isDecreasing = true
	for i := countDigits(num); i > 1; i-- {
		if returnDigit(num, i) < returnDigit(num, i-1) {
			isDecreasing = false
		} else {
			if returnDigit(num, i) > returnDigit(num, i-1) {
				isIncreasing = false
			}
		}
		if !(isIncreasing || isDecreasing) {
			return true
		}
	}
	return false
}

func findNumberByBouncyPercentage(percentage float64) (num int) {
	num = 1
	var bouncyCount = 0
	for {
		if isBouncy(num) {
			bouncyCount++
		}
		var proportion = float64(bouncyCount) / float64(num)
		if proportion == percentage {
			return num
		} else if proportion > percentage || percentage >= 1 {
			fmt.Printf("no number at which the bouncy proportion is exactly %d could be found", percentage)
			return num
		}
		num++
	}
}
