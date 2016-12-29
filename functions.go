package main

import (
    "bufio"
    "fmt"
    "os"
    "net/url"
    "log"
    "strings"
)

// func input() (string) {
//     reader := bufio.NewReader(os.Stdin)
//     fmt.Print("URL: ")
//     Homepage, _ := reader.ReadString('\n')
//     if !(strings.Contains(Homepage, "://")){
// 	    Homepage = "http://" + Homepage
//     }
//     return Homepage
// }
//
// func get_domain_name(raw_url string) string {
//     var new_url string
//     scheme := "http://"
//     if !strings.Contains(raw_url, scheme) {
//         s := []string{scheme, raw_url}
//         new_url = strings.Join(s, "")
//         // log.Panicln("Need url scheme, i.e. http://")
//     } else {
//         new_url = raw_url
//     }
//     u, err := url.Parse(new_url)
//     fmt.Println("Scheme: ", u.Scheme)
//     fmt.Println("Host: ", u.Host)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
//     return u.Host
// }

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
