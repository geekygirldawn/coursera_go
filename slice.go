// Write a program which prompts the user to enter integers and stores the integers 
// in a sorted slice. The program should be written as a loop. Before entering the 
// loop, the program should create an empty integer slice of size (length) 3. During 
// each pass through the loop, the program prompts the user to enter an integer to be 
// added to the slice. The program adds the integer to the slice, sorts the slice, and 
// prints the contents of the slice in sorted order. The slice must grow in size to 
// accommodate any number of integers which the user decides to enter. The program should 
// only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

// Submit your source code for the program, “slice.go”.

package main

import (
    "fmt"
    "strings"
    "strconv"
    "sort"
)

func main(){

    var user_int string
    sli := make([]int, 3)

    x := false
    i := 0

    fmt.Printf("The initialized slice contains: ")
    fmt.Print(sli)

    for x == false{

        fmt.Printf("\nPlease enter one integer: ")
        fmt.Scan(&user_int)

        if (strings.HasPrefix(user_int, "X") || strings.HasPrefix(user_int, "x")) {

              // Exit condition
              x = true

        } else if valid_int, err := strconv.Atoi(user_int); err == nil {

              // Condition where valid input is added to slice

              if (len(sli) <  i+1){
                  sli = append(sli, valid_int)
              } else {
                  sli[0] = valid_int
              }

              fmt.Printf("Current slice contains: ")
              sort.Ints(sli)
              fmt.Print(sli)

        } else{

            i-- // decrements i for invalid input where nothing is added to slice.
            fmt.Printf("Invalid input. Input must be an integer or 'X' to exit")
        }

        i++

    }
}
