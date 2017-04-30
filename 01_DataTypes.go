package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)

    // Declare second integer, double, and String variables.
    var myInt uint64
    var myFloat float64
    var myStr string

    // Read and save an integer, double, and String to your variables.
    scanner.Scan()
    myInt, _ = strconv.ParseUint(scanner.Text(), 10, 64)
    scanner.Scan()
    myFloat, _ = strconv.ParseFloat(scanner.Text(), 64)
    scanner.Scan()
    myStr = scanner.Text()

    // Print the sum of both integer variables on a new line.
    fmt.Println(i + myInt)

    // Print the sum of the double variables on a new line.
    fmt.Println(strconv.FormatFloat(d + myFloat, 'f', 1, 64))

    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    fmt.Println(s + myStr)
}