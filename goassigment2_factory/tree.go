package main

type Tree struct {
	MapObject
}

func newTree() imapObject {
	return &Tree{
		MapObject: MapObject{
			name:     "Tree",
			isLiving: false,
		},
	}
}
