package main

import (
    "bufio"
    "fmt"
    "os"
    "net/url"
    "log"
    "strings"
)

func input() (string) {

    scan := bufio.NewScanner(os.Stdin)
    fmt.Print("Input URL: ")
    scan.Scan()
    Homepage := scan.Text()
    if !(strings.Contains(Homepage, "://")) {
	    Homepage = "http://" + Homepage
        if !(strings.Contains(string(Homepage[len(Homepage)-1]), "/")) {
            Homepage = Homepage + "/"
        }
	}
    return Homepage
}

func get_domain_name(rawurl string) string {
    u, err := url.Parse(rawurl)
    fmt.Println("Scheme: ", u.Scheme)
    fmt.Println("Host: ", u.Host)
	if err != nil {
	       log.Fatalln(err)
	}
    return u.Host
}
