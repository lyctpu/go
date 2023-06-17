package main

import "fmt"

func main() {
    x:="Go"
    fmt.Println("Hello World", x)
    fmt.Println(len("Hello World"))
    fmt.Println("Hello World"[1])
    fmt.Println("Hello " + "World")
}
// go run hello.go
