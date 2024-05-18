package main

import "fmt"

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func (c Circle) circumference() float64 {
	return 2 * PI * c.radius
}

func main() {

	c := Circle{radius: 100}

	fmt.Println(c.circumference())

}
