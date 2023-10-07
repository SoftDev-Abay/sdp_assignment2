package main

type Character struct {
	MapObject
}

func newCharacter() imapObject {
	return &Character{
		MapObject: MapObject{
			name:     "Character",
			isLiving: true,
		},
	}
}
