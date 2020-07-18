package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	area() float64
}

func info(s shape) {
	area := s.area()
	fmt.Println(area)
}

func main() {
	sq := square{
		length: 23.5,
	}

	ci := circle{
		radius: 5.7,
	}

	info(sq)
	info(ci)

}
