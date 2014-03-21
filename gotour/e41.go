package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    wordMap := make(map[string]int)
    for _, word := range strings.Fields(s) {
        wordMap[word] += 1
    }
    return wordMap
}

func main() {
    wc.Test(WordCount)
}