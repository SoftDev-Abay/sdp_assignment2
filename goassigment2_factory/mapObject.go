package main

type MapObject struct {
	name     string
	isLiving bool
}

func (mo *MapObject) setName(name string) {
	mo.name = name
}

func (mo *MapObject) getName() string {
	return mo.name
}

func (mo *MapObject) setIsLiving(isLiving bool) {
	mo.isLiving = isLiving
}

func (mo *MapObject) getIsLiving() bool {
	return mo.isLiving
}
