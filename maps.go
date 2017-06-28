package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["z1"] = 11
    m["z14"] = 14

    fmt.Println(m)

    // get value of key exist, otherwise return "default" value
    _, exist := m["z3"]  // _ ignores the returned value; can't be used as variable name
    fmt.Println(exist)
    // _ => blank identifier

    // map length
    fmt.Println(len(m))

    // one-liner
    z := map[string]int{"one": 1, "two": 2}
    fmt.Println(z)
}
