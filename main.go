package go_crawl

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var VERSION string = "1.0"
var ProjectName string

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func check(e error) {
	if e != nil {
		fmt.Println("*********************ERROR*******************")
		fmt.Println(e)
		panic(e)
	}
}

func create_dirs() {
	var path bytes.Buffer
	path.WriteString("projects" + string(filepath.Separator) + ProjectName)
	fmt.Println("The path is !!!!XXXX: " + path.String())
	os.MkdirAll(path.String(), os.ModePerm)
	// check(e)
	queue_file := filepath.Join(path.String(), "queue.txt")
	crawled_file := filepath.Join(path.String(), "crawled.txt")
	errors_file := filepath.Join(path.String(), "errors.txt")
	summary_file := filepath.Join(path.String(), "summary.txt")
	a := [5]string{queue_file, crawled_file, errors_file, summary_file}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		f, _ := os.Create(a[i])
		// check(e)
		f.Close()
	}
}

func getDomainName(rawurl string) string {
	u, err := url.Parse(rawurl)
	fmt.Println("Scheme: ", u.Scheme)
	fmt.Println("Host: ", u.Host)
	check(err)
	return u.Host
}

func Crawl(url string, threads int) {
	fmt.Println("go_crawl Version: ", VERSION)
	DomainName := getDomainName(url)
	fmt.Println("URL: ", DomainName)
	ProjectName = DomainName[:strings.Index(DomainName, ".")]
	fmt.Println(ProjectName)
	create_dirs()
}
