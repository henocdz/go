package main

import "errors"
import "fmt"

func noError(number int) (int, error) {
    if (number % 2 != 0) {
        return 0, errors.New("don't work with odd numbers")
    }
    return number * number, nil
}

type CustomError struct {
    text string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("%s", e.text)
}

func customError() (int, error) {
    return 0, &CustomError{"didn't work"}
}

func main() {
    num, err := noError(4)
    fmt.Println(num, err)

    if _, e := customError(); e != nil {
        // print value returned by `CustomError.Error` method
        fmt.Println(e)
        // access to variables defined in struct CustomError
        // via type assertion: https://tour.golang.org/methods/15
        if err, ok := e.(*CustomError); ok {
            fmt.Println(err.text)
        }
    }
    // errors are values by Rob Pike: https://blog.golang.org/errors-are-values
}
