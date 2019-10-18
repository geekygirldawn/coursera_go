package main

import (
	"fmt"
)

func race (a string){

    fmt.Println(a)

}

func main(){

    for i := 0; i < 5; i++ {
        go race("Hello")
        race("World")
    }
}


/*
The program has a simple function that prints a string.
This is called from within a loop as a goroutine:
go race("Hello")

Along with a second call to the same function not as a goroutine:
race("World")

The regular function always prints "World" during every execution 
of the loop, but the "Hello" printed by the goroutine appears at
unreliable intervals, thus containing a race condition.

As you can see, the output is different each time the program is run,
indicating the race condition.

Dawns-MacBook-Pro:concurrency dafoster$ go run race.go 
Hello
World
World
World
Hello
World
World
Dawns-MacBook-Pro:concurrency dafoster$ go run race.go 
World
World
Hello
World
World
World
Hello
Hello
Hello
Dawns-MacBook-Pro:concurrency dafoster$ go run race.go 
World
World
Hello
World
*/

