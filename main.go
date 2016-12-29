package main

import (
    "fmt"
)

func main() {
    url := input()
    Domain_name := get_domain_name(url)
    fmt.Println("Domain Name: ", Domain_name)
}
