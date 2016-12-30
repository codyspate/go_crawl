package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	//"io/ioutil"
)

func check(e error) {
	if e != nil {
		fmt.Println("*********************ERROR*******************")
		fmt.Println(e)
		panic(e)
	}
}

func input() string {
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
	check(err)
	return u.Host
}

func create_dirs() {
	var path bytes.Buffer
	path.WriteString("projects" + string(filepath.Separator) + project_name)
	fmt.Println("The path is !!!!XXXX:   " + path.String())
	e := os.MkdirAll(path.String(), os.ModePerm)
	check(e)
	queue_file = filepath.Join(path.String(), "queue.txt")
	crawled_file = filepath.Join(path.String(), "crawled.txt")
	errors_file = filepath.Join(path.String(), "errors.txt")
	summary_file = filepath.Join(path.String(), "summary.txt")
	a := [5]string{queue_file, crawled_file, errors_file, summary_file}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		f, e := os.Create(a[i])
		check(e)
		f.Close()
	}
}
