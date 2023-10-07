package main

type House struct {
	MapObject
}

func newHouse() imapObject {
	return &House{
		MapObject: MapObject{
			name:     "House",
			isLiving: false,
		},
	}
}
