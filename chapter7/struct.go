package main

import (
	"fmt"
	"math"
)

func main() {

	c := Circle{1, 1, 1}

	r := Rectangle{0, 0, 10, 10}

	fmt.Println("Circle's area", c.area())
	fmt.Println("Rectangles's area", r.area())
	fmt.Println("Total area", totalArea(&c, &r))

	fmt.Println("Perimeter of the circle", c.perimeter())
	fmt.Println("Perimeter of the rectangle", r.perimeter())

	fmt.Println("Total perimeter ", totalPerimeter(&c, &r))
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {

	c.r = 4
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	x := x2 - x1
	y := y2 - y1

	return math.Sqrt(x*x + y*y)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {

	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w

}

func (r Rectangle) perimeter() float64 {

	return r.x1 + r.x2 + r.y1 + r.y2
}

func (c Circle) perimeter() float64 {

	return c.x + c.y
}

func totalArea(shapes ...Shape) float64 {

	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.area()
	}

	return totalArea
}

func totalPerimeter(shapes ...Shape) float64 {

	totalPerimeter := 0.0
	for _, s := range shapes {
		totalPerimeter += s.perimeter()
	}

	return totalPerimeter
}

type Shape interface {
	area() float64
	perimeter() float64
}
