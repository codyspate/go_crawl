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
    return Homepage
}

func get_domain_name(rawurl string) string {
    scheme := "http://"
    if strings.Contains(rawurl, scheme) == false {
        log.Panicln("Need url scheme, i.e. http://")
    }
    u, err := url.Parse(rawurl)
    fmt.Println("Scheme: ", u.Scheme)
    fmt.Println("Host: ", u.Host)
	if err != nil {
		log.Fatalln(err)
	}
    return u.Host
}
