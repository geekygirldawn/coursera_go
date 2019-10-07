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

type Animal struct {
    food, locomotion, noise string
}

func (a Animal) Eat() {

   fmt.Println(a.food)

}

func (a Animal) Move() {

   fmt.Println(a.locomotion)

}

func (a Animal) Speak() {

   fmt.Println(a.noise)

}

func main(){

    cow := Animal{
        food: "grass",
        locomotion: "walk",
        noise: "moo",
    }
    bird := Animal{
        food: "worms",
        locomotion: "fly",
        noise: "peep",
    }
    snake := Animal{
        food: "mice",
        locomotion: "slither",
        noise: "hsss",
    }

    for {
        invalid := 0

        words := read_input()
        a := words[0]
        i := words[1]

        if a == "cow"{
            if i == "speak"{
                cow.Speak()
            } else if i == "move"{
                cow.Move()
            } else if i == "eat"{
                cow.Eat()
            } else {
                invalid = 1
            }
        } else if a == "bird"{
            if i == "speak"{
                bird.Speak()
            } else if i == "move"{
                bird.Move()
            } else if i == "eat"{
                bird.Eat()
            } else {
               invalid = 1
            }
        } else if a == "snake"{
            if i == "speak"{
                snake.Speak()
            } else if i == "move"{
                snake.Move()
            } else if i == "eat"{
                snake.Eat()
            } else {
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
