package main

import (
    "fmt"

    "github.com/wmy09527/mytool/greetings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
