package main

import "fmt"

type Chicken struct {
	Name string
}

func (c Chicken) Eat() {
	fmt.Println(c.Name, "is eating")
}
func (c Chicken) Move() {
	fmt.Println(c.Name, "is moving")
}
func (c Chicken) Speak() {
	fmt.Println(c.Name, "is speaking")
}
