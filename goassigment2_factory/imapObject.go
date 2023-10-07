package main

type imapObject interface {
	setName(name string)
	setIsLiving(isLiving bool)
	getName() string
	getIsLiving() bool
}
