package main

import (
	"fmt"
)

func createMap(treeCount int, houseCount int, characterCount int) Map {
	mapObjects := []imapObject{}

	for i := 0; i < treeCount; i++ {
		tree, err := createMapObject("Tree")
		if err != nil {
			fmt.Println(err)
			continue
		}
		mapObjects = append(mapObjects, tree)
	}

	for i := 0; i < houseCount; i++ {
		house, err := createMapObject("House")
		if err != nil {
			fmt.Println(err)
			continue
		}
		mapObjects = append(mapObjects, house)
	}

	for i := 0; i < characterCount; i++ {
		character, err := createMapObject("Character")
		if err != nil {
			fmt.Println(err)
			continue
		}
		mapObjects = append(mapObjects, character)
	}

	m := &Map{
		objects: mapObjects,
	}
	return *m
}
