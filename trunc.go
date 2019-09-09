// ASSIGNMENT: Module 2
// Write a program which prompts the user to enter a floating point number and 
// prints the integer which is a truncated version of the floating point number 
// that was entered. Truncation is the process of removing the digits to the 
// right of the decimal place.

// Submit your source code for the program, “trunc.go”.

package main

import "fmt"

func main(){

    var fp_num float64

    fmt.Printf("Please enter a floating point number: ")
    num, err := fmt.Scan(&fp_num)

    trunc := int(fp_num)

    if (num == 1 && err == nil){
        fmt.Printf("Your integer is: %d\n", trunc)
    } else{
        fmt.Printf("Input Error\n")
    }
}

