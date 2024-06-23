package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}

    fmt.Println("Elemen-elemen slice:")
    for _, num := range numbers {
        fmt.Println(num)
    }
}
