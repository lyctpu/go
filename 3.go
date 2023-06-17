package main

import "fmt"

func main() {
  var total float64 = 30
  x := [5]int{ 98, 93, 77, 82, 83 }
  fmt.Println(total / float64(len(x)))

  x_srez:= []int{ 98, 93, 77, 82, 83 }
  x_srez2 := make([]int, 2)
  copy(x_srez2, x_srez1)

  sr := x[0:5]
  slice2 := append(sr, 4, 5)
  fmt.Println(sr, slice2)

}
