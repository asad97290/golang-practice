package main

import "fmt"

// this is a comment
func fact(num int) int {
    if num == 0 {
        return 1
    }
    return num * fact(num-1)
}
func main() {

    var n int = 3
    const num = 3.25
    number := n
    fmt.Println("Factorial of",number,"is",fact(n))


    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }

    i := 4

    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
        
    default:
        fmt.Println("default")

    }
}