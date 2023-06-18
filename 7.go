package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}



func main() {
    var rx1, ry1 float64 = 0, 0
    var rx2, ry2 float64 = 10, 10
    var cx, cy, cr float64 = 0, 0, 5
  
    r := Rectangle{0, 0, 10, 10}
    fmt.Println(r.area())
  
    x := 5
    zero(&x)
    fmt.Println(x) // x is 0
    
    xPtr := new(int)
    one(xPtr)
    fmt.Println(*xPtr) // x is 1
  
    //c := new(Circle)
  //c := Circle{x: 0, y: 0, r: 5}
    c := Circle{0, 0, 5}
    fmt.Println(c.x, c.y, c.r)
    c.x = 10
    c.y = 5
    fmt.Println(circleArea(&c))
    fmt.Println(c.area())
  
    a := new(Android)
    a.Talk()
  
    fmt.Println(totalArea(&c, &r))
  
}

func zero(xPtr *int) {
    *xPtr = 0
}

func one(xPtr *int) {
    *xPtr = 1
}

type Circle struct {
    x, y, r float64
}

func circleArea(c Circle) float64 {
    return math.Pi * c.r*c.r
}

func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
    Person
    Model string
}

type Shape interface {
    area() float64
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

type MultiShape struct {
    shapes []Shape
}

func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}
