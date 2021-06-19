package main

import "fmt"

func main() {
	slice := generateSliceOfSquentialInt(1, 10)
	for _, number := range slice {
		if isEven(number) {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}
}

func generateSliceOfSquentialInt(startNum int, endNum int) []int {
	slice := []int{}
	for i := startNum; i <= endNum; i++ {
		slice = append(slice, i)
	}
	return slice
}

func isEven(number int) bool {
	return number%2 == 0
}
