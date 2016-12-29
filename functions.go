package main

import (
    "bufio"
    "fmt"
    "os"
)

func input() (string) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("URL: ")
    Homepage, _ := reader.ReadString('\n')
    return Homepage
}
