package main
import "fmt"

func main() {
    var N int
    fmt.Scan(&N)
    var A = make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&A[i])
    }
    
    for i := len(A) - 1; i >= 0; i-- {
        fmt.Print(A[i], " ")
    }
}
