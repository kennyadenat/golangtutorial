package main

import ("fmt"
"math")

func main() {

}

type Shape interface {
  area() float64
}

type Rectangle struct {
  height float64
  width float64
}