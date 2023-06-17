package main

import "fmt"

func main() {
    var b string
    b = "Hello World"
    
    var c string = "hello"
    const s string = "Hello World"
    
    x:="Go"
    fmt.Println("Hello World", x)
    fmt.Println(len("Hello World"))
    fmt.Println("Hello World"[1])
    fmt.Println("Hello " + "World")
}

func logic{
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}
// go run hello.go
