package main

import ("fmt"; "math")

func main () {

  c := Circle {1, 1 , 1}

  fmt.Println(c.area())

  fmt.Println(c)
}

type Circle struct{
  x float64
  y float64
  r float64
}

func (c *Circle) area() float64 {

  c.r = 4
  return math.Pi * c.r * c.r
}
