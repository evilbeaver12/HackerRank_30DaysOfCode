package main

import (
    "fmt"
    "strconv"
)

func main() {
    var mealCost, tipPercent, taxPercent float64
    fmt.Scan(&mealCost, &tipPercent, &taxPercent)
    totalCost := mealCost * (1 + (tipPercent + taxPercent) / 100)
    fmt.Println("The total meal cost is " + strconv.Itoa(int(totalCost + 0.5)) + " dollars.")
}
