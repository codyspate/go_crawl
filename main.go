package go_crawl

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var version = "1.0"
var projectName string
var crawled []string
var links []string

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

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println(err)
	return true
}

func writeFile(path string, text string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// write some text to file
	_, err = file.WriteString(text)
	checkError(err)
	_, err = file.WriteString(text)
	checkError(err)

	// save changes
	err = file.Sync()
	checkError(err)
}

func createDirs() {
	var path bytes.Buffer
	path.WriteString("projects" + string(filepath.Separator) + projectName)
	fmt.Println("The path is !!!!XXXX: " + path.String())
	os.MkdirAll(path.String(), os.ModePerm)
	// check(e)
	queueFile := filepath.Join(path.String(), "queue.txt")
	crawledFile := filepath.Join(path.String(), "crawled.txt")
	errorsFile := filepath.Join(path.String(), "errors.txt")
	summaryFile := filepath.Join(path.String(), "summary.txt")
	a := [5]string{queueFile, crawledFile, errorsFile, summaryFile}
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

func addLinksToQueue(links []string) {
	for key, value := range links {
		fmt.Println(key)
		fmt.Println(value)
	}
	// for url in links:
	//     if (url in Spider.queue) or (url in Spider.crawled):
	//         continue
	//     if not Spider.domain_name in get_domain_name(url):
	//         # try:
	//         #     ext_link = quote(url, safe="/:()=?#%&")
	//         #     response = urlopen(ext_link, timeout=10)
	//         #     response.close()
	//         # except Exception as e:
	//         #     Spider.num_broken_ext_links += 1
	//         #     append_to_file(Spider.ext_link_errors_file, url + " : " + str(e))
	//         continue
	//     url = quote(url, safe="%/:=&?~+!$,;'@()*[]#")
	//     if '%20' in url[-3:]:
	//         url = url[:-3]
	//     Spider.queue.add(url)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func crawlPage(pageURL string) {
	if !stringInSlice(pageURL, crawled) {
		fmt.Println("Now crawling: ", pageURL)

		apppend(crawled, pageURL)
	}
}

/*
Crawl crawls a website at the given url and will start
a number of concurrent processes depending on the threads input

Params:
	url <string> -----
		http address to website, can be IP or domain name, Ex. 71stsog.com, 12.345.45.67, http://google.com/

	threads <int> ----
		integer to specifiy the amount of concurrent processes to run
*/
func Crawl(url string, threads int) {
	fmt.Println("go_crawl Version: ", version)
	DomainName := getDomainName(url)
	fmt.Println("URL: ", DomainName)
	projectName = DomainName[:strings.Index(DomainName, ".")]
	fmt.Println(projectName)
	createDirs()
	crawlPage(url)
	for i := 0; i <= threads; i++ {
		go crawlPage(links[0])
	}
}
