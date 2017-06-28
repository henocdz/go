package main

import "fmt"

func main() {
    // python-like for k,v in dict.items()
    for k, v := range map[string]int{"one": 1, "two": 2} {
        fmt.Printf("%s: %d \n", k, v)
    }

    // more about how strings works, what is as rune, char, etc: https://blog.golang.org/strings
    for i, runeVal := range "❤☀☆☂☻♞☯☭☢€¡yeah!" {
        fmt.Printf("%#U at %d\n", runeVal, i)
    }
}
