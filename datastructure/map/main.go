package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":  "#FF0000",
		"blue": "#0000FF",
	}
	printMap(colors) // Key: red Value: #FF0000
	//  Key: blue Value: #0000FF

	colors2 := make(map[int]string)
	colors2[99] = "#FFFFFF"
	fmt.Println(colors2) // map[99:#FFFFFF]
	delete(colors2, 99)
	fmt.Println(colors2) // map[]

}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Key: " + key + " Value: " + value)
	}

}
