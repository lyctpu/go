//Горутина — это функция, которая может работать параллельно с другими функциями. 

package main

import (
    "fmt"
    "time"
    "math/rand"
)

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    }
}

func main() {
    for i := 0; i < 10; i++ {
        go f(i)
    }
    var input string
    fmt.Scanln(&input)
  
    var c chan string = make(chan string)

    go pinger(c)
    go ponger(c)
    go printer(c)
  
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for {
            c1 <- "from 1"
            time.Sleep(time.Second * 2)
        }
    }()
    go func() {
        for {
            c2 <- "from 2"
            time.Sleep(time.Second * 3)
        }
    }()
    go func() {
        for {
            select {
            case msg1 := <- c1:
                fmt.Println(msg1)
            case msg2 := <- c2:
                fmt.Println(msg2)
            }
        }
    }()
  
  
  
  
  
  
  
  
  
}

func pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping"
    }
}

func ponger(c chan string) {
    for i := 0; ; i++ {
        c <- "pong"
    }
}

func printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}

Программа будет постоянно выводить «ping» 
(нажмите enter, чтобы её остановить). 
Тип канала представлен ключевым словом chan, 
за которым следует тип, который будет передаваться по каналу 
(в данном случае мы передаем строки). 
Оператор <- (стрелка влево) используется для отправки и получения сообщений по каналу. 
Конструкция c <- "ping" означает отправку "ping", 
а msg := <- c — его получение и сохранение в переменную msg. 
Строка с fmt может быть записана другим способом: fmt.Println(<-c), 
тогда можно было бы удалить предыдущую строку.
Данное использование каналов позволяет синхронизировать две горутины. 
Когда pinger пытается послать сообщение в канал, 
он ожидает, пока printer будет готов получить сообщение. 
Такое поведение называется блокирующим.

Эта программа выводит «from 1» каждые 2 секунды и «from 2» каждые 3 секунды. 
Оператор select выбирает первый готовый канал, 
и получает сообщение из него, или же передает сообщение через него. 
Когда готовы несколько каналов, получение сообщения происходит 
из случайно выбранного готового канала. Если же ни один из каналов не готов, 
оператор блокирует ход программы до тех пор, пока какой-либо 
из каналов будет готов к отправке или получению.
