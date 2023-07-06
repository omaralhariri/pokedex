package main

import (
    "fmt"
    "os"
)

func main() {
    acceptCommands();
}

func acceptCommands() {
    fmt.Println("Please enter a command")

    args := os.Args[1:]

    for _, v := range args {
        fmt.Println(v)
    }
}
