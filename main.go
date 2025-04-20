package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	pers := person{"Vishnu", 90}
	fmt.Println(pers.getNmae())
}

type person struct {
	name string
	age  int
}

func (p person) getNmae() string {
	return p.name
}
