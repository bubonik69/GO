package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age int
	Type string
}

type Dog struct {
	Animal
}

type Voiced interface {
	Voice()
}

func (d *Animal) Voice() {
	fmt.Println("Myyyy")
}

func (d Animal) incrementage() int {
	d.Age+=1
	return d.Age
}
func (d *Animal) getAge() {
	fmt.Println(d.Age)
}
func (d *Dog) Voice() {
	fmt.Println("gav")
}

func (d *Cat) Voice() {
	fmt.Println("mya")
}

type Cat struct {
	Animal
}

func main() {
	//dog := Dog{}
	//cat := Cat{}
	//dog.Voice()
	//cat.Voice()
	animal:=Animal{}
	animal.Age=animal.incrementage()
	animal.getAge()
	animal.incrementage()
	animal.getAge()
	animal.incrementage()
	animal.getAge()
}