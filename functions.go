package main

import "fmt"

func add(x , y int) (bob int) {
    bob = x + y
    return bob
}

func main() {
    fmt.Println(add(42,13))
}