package main

import "fmt"

func first() {
    fmt.Println("1st")
    f, _ := os.Open(filename)
    defer f.Close()
}
func second() {
    fmt.Println("2nd")
}
func main() {
    defer second()
    first()
  
    defer func() {    
       str := recover()
       fmt.Println(str)
    }()
    panic("PANIC")
}

//Такой подход дает нам три преимущества: 
//1) вызовы Close и Open располагаются рядом, что облегчает понимание программы, 
//(2) если функция содержит несколько операций возврата 
//(например, одна произойдет в блоке if, другая в блоке else), 
//Close будет вызван до выхода из функции, 
//(3) отложенные функции вызываются, даже если во время выполнения происходит ошибка.

