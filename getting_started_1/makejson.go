// Write a program which prompts the user to first enter a name, and then enter an address. 
// Your program should create a map and add the name and address to the map using the keys “name” 
// and “address”, respectively. Your program should use Marshal() to create a JSON object from the 
// map, and then your program should print the JSON object.

// Submit your source code for the program, “makejson.go”.

package main

import (
    "fmt"
    "bufio"
    "os"
    "encoding/json"
)

func main(){

    var name string
    var addr string

    person := make(map[string]string)

    read_input := bufio.NewScanner(os.Stdin)

    fmt.Printf("\nPlease enter a name: ")
    read_input.Scan()
    name = read_input.Text()

    fmt.Printf("Please enter an address: ")
    read_input.Scan()
    addr = read_input.Text()

    person["name"] = name
    person["address"] = addr

    jsonperson, _ := json.Marshal(person)
    fmt.Println("Your JSON Object is: ")
    fmt.Println(string(jsonperson))
}
