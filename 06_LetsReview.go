package main
import "fmt"

func main() {
    var T int
    fmt.Scan(&T)
    for i:=0; i < T; i++ {
        var S string
        fmt.Scan(&S)
        evenS := ""
        oddS := ""
        for j := 0; j < len(S); j++ {
            if j % 2 == 0 {
                evenS += S[j:j+1]
            } else {
                oddS += S[j:j+1]
            }
        }
        fmt.Println(evenS, oddS)
    }
}
