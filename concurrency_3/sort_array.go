/******** 
Write a program to sort an array of integers. The program should partition the array 
into 4 parts, each of which is sorted by a different goroutine. Each partition 
should be of approximately equal size. Then the main goroutine should merge the 4 
sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine 
which sorts ¼ of the array should print the subarray that it will sort. When 
sorting is complete, the main goroutine should print the entire sorted list. 
**********/

package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "sort"
)

func read_ints() []int{

    // Reads space separated list of integers from stdin, validates input,
    // and stores it in a slice. 

    read_input := bufio.NewScanner(os.Stdin)
    fmt.Println("\nPlease enter a sequence of integers separated by spaces.")
    fmt.Println("Hit enter when finished:")
    read_input.Scan()
    line := read_input.Text()

    user_ints := strings.Split(line, " ")

    var user_slice []int

    for i:=0; i < len(user_ints); i++ {
        user_int, err := strconv.Atoi(user_ints[i])
        if err == nil {
            user_slice = append(user_slice, user_int)
        } else {
            fmt.Println("Invalid input. Please run the program again with ")
            fmt.Println("a sequence of integers separated by spaces. ")
            fmt.Print("Exiting program ")
            os.Exit(1)
        }

    }
    return user_slice

}

func check_rem(rem int) (int, int) {

    inc := 0
    if rem > 0{
        inc = 1
        rem = rem - 1
    }
    return inc, rem
}

func split_input(user_slice []int) ([]int, []int, []int, []int){

    // Takes the slice of integers and splits into 4 arrays of approx equal size
    num_ints := len(user_slice)
    chunks := 4
    per_slice :=  num_ints / chunks
    remainder := num_ints % chunks

    var inc1, rem1 = check_rem(remainder)
    array1 := user_slice[:per_slice+inc1]
    var inc2, rem2 = check_rem(rem1)
    array2 := user_slice[per_slice+inc1:per_slice*2+inc1+inc2]
    var inc3, _ = check_rem(rem2)
    array3 := user_slice[per_slice*2+inc1+inc2:per_slice*3+inc1+inc2+inc3]
    array4 := user_slice[per_slice*3+inc1+inc2+inc3:per_slice*4+inc1+inc2+inc3]

    return array1, array2, array3, array4
}

func sort_array(arr []int, c chan[]int){

    orig_arr := make([]int, len(arr))
    copy(orig_arr, arr)
    sort.Ints(arr)
    fmt.Println("Sorted", orig_arr, "into", arr)

    c <- arr

}

func main(){

    user_slice := read_ints()

    array1, array2, array3, array4 := split_input(user_slice)

    fmt.Println("\nSorting partitions ...")

    c := make(chan []int, 4)
    go sort_array(array1, c)
    go sort_array(array2, c)
    go sort_array(array3, c)
    go sort_array(array4, c)
    sort1 := <-c
    sort2 := <-c
    sort3 := <-c
    sort4 := <-c

    var merged []int
    merged = append(merged, sort1...)
    merged = append(merged, sort2...)
    merged = append(merged, sort3...)
    merged = append(merged, sort4...)
    sort.Ints(merged)
    fmt.Println("\nThe final sort is:")
    fmt.Println(merged)
}
