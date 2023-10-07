package main

import "fmt"

func createMapObject(mapObjectType string) (imapObject, error) {
	if mapObjectType == "Tree" {
		return newTree(), nil
	}
	if mapObjectType == "House" {
		return newHouse(), nil
	}
	if mapObjectType == "Character" {
		return newCharacter(), nil
	}
	return nil, fmt.Errorf("Wrong map object type passed")
}
