package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func read_input() []string {

    read_input := bufio.NewScanner(os.Stdin)
    fmt.Printf(">")

    read_input.Scan()
    temp := read_input.Text()
    words := strings.Fields(temp)
    return words
}

type Animal interface {
    Eat()
    Move()
    Speak()
}

type Cow struct {name string}
func (c Cow) Eat(){
    fmt.Println("grass")
}
func (c Cow) Move(){
    fmt.Println("walk")
}
func (c Cow) Speak(){
    fmt.Println("moo")
}

type Bird struct {name string}
func (b Bird) Eat(){
    fmt.Println("worms")
}
func (b Bird) Move(){
    fmt.Println("fly")
}
func (b Bird) Speak(){
    fmt.Println("peep")
}

type Snake struct {name string}
func (s Snake) Eat(){
    fmt.Println("mice")
}
func (s Snake) Move(){
    fmt.Println("slither")
}
func (s Snake) Speak(){
    fmt.Println("hsss")
}

func main(){

    animalMap := map[string][]string{}

    for {
        invalid := 0
        var command string
        var name string
        var info string

        words := read_input()
        if len(words) == 3{
            command = words[0]
            name = words[1]
            info = words[2]
        } else {
            invalid = 1
        }

        if command == "newanimal"{
            animalMap[name] = append(animalMap[name], info)
            fmt.Println("Created it!")
        } else if command == "query"{
            var a Animal
            if animalMap[name][0] == "cow"{
                c := Cow{name:name}
                a = c
            } else if animalMap[name][0] == "snake"{
                s := Snake{name:name}
                a = s
            } else if animalMap[name][0] == "bird"{
                b := Bird{name:name}
                a = b
            } else {
            invalid = 1
            }

            if info == "speak"{
                a.Speak()
            } else if info == "eat" {
                a.Eat()
            } else if info == "move" {
                a.Move()
            } else{
                invalid = 1
            }

        } else {
            invalid = 1
        }

        if invalid == 1{
            fmt.Println("Invalid input. Please try again.")
        }
    }
}
