// Write a program which prompts the user to enter a string. The program searches 
// through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program 
// should print “Found!” if the entered string starts with the character ‘i’, 
// ends with the character ‘n’, and contains the character ‘a’. The program 
// should print “Not Found!” otherwise. The program should not be case-sensitive, 
// so it does not matter if the characters are upper-case or lower-case.

// Examples: The program should print “Found!” for the following example entered 
// strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should 
// print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

// Submit your source code for the program, “findian.go”.

package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func main(){

    var ian string

    fmt.Printf("Please enter one string: ")
    input := bufio.NewReader(os.Stdin)
    ian_untrimmed, err := input.ReadString('\n')

    if (err == nil){
        ian = strings.TrimSuffix(ian_untrimmed, "\n")
        ian_lower := strings.ToLower(ian)

        starts := strings.HasPrefix(ian_lower, "i")
        ends := strings.HasSuffix(ian_lower, "n")
        contains := strings.Contains(ian_lower, "a")

        if (starts && contains && ends){

            fmt.Printf("Found!\n")
        } else{
            fmt.Printf("Not Found!\n")
        }
    } else{
        fmt.Printf("Input Error\n")
    }
}
