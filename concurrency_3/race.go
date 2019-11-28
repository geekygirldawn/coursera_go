package main

import (
    "fmt"
    "time"
)

func race (a string){

    fmt.Println(a)

}

func main(){

    for i := 0; i < 5; i++ {
        go race("Hello")
        go race("World")
        time.Sleep(10 * time.Millisecond)
    }
}


/*
The program has a simple function that prints a string.
This is called from within a loop as 2 goroutines:
go race("Hello")
go race("World")

"Hello" and "World" printed by the goroutines appear at
unreliable intervals, thus containing a race condition.

As you can see, the output is different each time the program is run,
indicating the race condition.

Dawns-MacBook-Pro:concurrency dafoster$ go run race.go 
Hello
World
World
Hello
World
Hello
Hello
World
World
Hello
Dawns-MacBook-Pro:concurrency dafoster$ go run race.go 
World
Hello
Hello
World
Hello
World
Hello
World
Hello
World
Dawns-MacBook-Pro:concurrency dafoster$ go run race.go 
World
Hello
World
Hello
Hello
World
Hello
World
Hello
World
*/

