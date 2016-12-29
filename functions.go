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
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("URL: ")
    Homepage, _ := reader.ReadString('\n')
    if !(strings.contains(Homepage, "://")){
	    Homepage = "http://" + Homepage
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
