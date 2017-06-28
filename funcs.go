package main

import "fmt"

func balls() (int, int, string) {
    return 0, 0, "third"
}

func sumBy(nums ...int, divisors ...int) (float64) {
    var total float64
    for index, num := range nums {
        divisor := divisors[index]
        total += float64(num / divisor)
    }
    return total
}

func main() {
    ballOne, ballTwo, _ := balls()
    fmt.Println(ballOne, ballTwo)

    _, _, third := balls()  // ignore returned values
    fmt.Println(third)
}
