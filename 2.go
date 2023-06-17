package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)
    output := input * 2
    fmt.Println(output)
  
    fmt.Println(`1
2
3
`)
  
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i = i + 1
    }
  
  for i := 1; i <= 10; i++ {
        fmt.Println(i)
        if i % 2 == 0 {
            // divisible by 2
        } else if i % 3 == 0 {
            // divisible by 3
        } else if i % 5 == 0 {
            // divisible by 5
        }
   }
  
  switch i {
  case 0: fmt.Println("Zero")
  case 1: fmt.Println("One")
  case 2: fmt.Println("Two")
  case 3: fmt.Println("Three")
  case 4: fmt.Println("Four")
  case 5: fmt.Println("Five")
  default: fmt.Println("Unknown Number")
  }
  
  
}
