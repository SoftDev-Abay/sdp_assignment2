package main

import (
	"fmt"
)

func main() {
	fmt.Println("Creating a map!")
	fmt.Println("How many trees?")
	var treeCount int
	fmt.Scanln(&treeCount)

	fmt.Println("How many houses?")
	var houseCount int
	fmt.Scanln(&houseCount)

	fmt.Println("How many characters?")
	var characterCount int
	fmt.Scanln(&characterCount)

	m := createMap(treeCount, houseCount, characterCount)
	fmt.Println("Map created!")
	fmt.Println("Map objects:")
	for _, obj := range m.GetObjects() {
		fmt.Println(obj.getName())
	}
}
