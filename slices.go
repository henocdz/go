package main

import "fmt"


func main() {
    z := []string{}
    z = append(z, "3", "4")
    fmt.Println(z)
    // if c slice is created with len-0 ...
    c := make([]string, len(z))
    // ... this will not copy anything since doesn't have where to store copies
    copy(c, z)
    fmt.Println(c)
    fmt.Println("slice-c", c[:1])
}
