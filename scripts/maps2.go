package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    wordc := make(map[string]int)
    sentence := strings.Fields(s)

    for _, v := range sentence {
        wordc[v] += 1
    }
    return wordc
}

func main() {
    wc.Test(WordCount)
}