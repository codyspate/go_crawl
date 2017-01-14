package go_crawl

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"unicode/utf8"

	"golang.org/x/net/html"
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

func parseHTMLforAnchorTags(r io.Reader) []string {
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
	// Parse HTML for Anchor Tags //
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
	fmt.Println("Parsing html for anchor tags")
	page := html.NewTokenizer(r)
	fmt.Println("page: ", page)
	fmt.Println("page.Raw(): ", page.Raw())
	fmt.Println("string(page.Raw())", string(page.Raw()))
	fmt.Println("Err: ", page.Err())
	fmt.Println("Token().String(): ", page.Token().String())
	fmt.Println("Token().Data: ", page.Token().Data)
	fmt.Println("string(page.Text()): ", string(page.Text()))
	fmt.Println("page.Text(): ", page.Text())
	for {
		fmt.Println("Starting for loop")
		fmt.Println(page.Text())
		tokenType := page.Next()
		fmt.Println("tokeType: ", tokenType)
		if tokenType == html.ErrorToken {
			fmt.Println("ErrorToken")
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
					// tl := trimHash(attr.Val)
					// col = append(col, tl)
					// resolv(&links, col)
				}
			}
		}
		// switch {
		// case tt == html.ErrorToken:
		// 	fmt.Println("End of the document, we're done...")
		// 	// End of the document, we're done
		// 	return
		// case tt == html.StartTagToken:
		// 	t := z.Token()

		// 	if t.Data == "a" { // is anchor
		// 		fmt.Println("We found a link!")
		// 		fmt.Println(t)
		// 		fmt.Println(t.Data)
		// 	}
		// default:
		// 	fmt.Println(tt)
		// 	fmt.Println(z.Token())
		// 	fmt.Println(string(z.Token().Data))
		// 	fmt.Println(string(z.Raw()))
		// }
	}
}

// Author: Dixon Begay
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

// Authors: Timmothy Spate, Dixon Begay
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

// Author: Dixon Begay
func getProperURL(rawurl string) *url.URL {
	fmt.Println("Inputed URL: ", rawurl)
	if !strings.Contains(rawurl, "http://") && !strings.Contains(rawurl, "https://") {
		// fmt.Println("No http header... Adding to url")
		rawurl = fmt.Sprintf("http://" + rawurl)
		// fmt.Println(rawurl)
	}
	if !(strings.Contains(rawurl, ".com")) {
		// fmt.Println("No '.com' in url... Adding to url")
		rawurl = fmt.Sprintf(rawurl + ".com")
		// fmt.Println(rawurl)
	}
	if rawurl[len(rawurl)-1:] != "/" {
		// fmt.Println("No trailing forward slash... Adding to url")
		rawurl = fmt.Sprintf(rawurl + "/")
		// fmt.Println(rawurl)
	}
	// fmt.Println("Using this url to parse: ", rawurl)
	u, err := url.Parse(rawurl)
	fmt.Println("Scheme: ", u.Scheme)
	fmt.Println("Host: ", u.Host)
	check(err)
	return u
}

// Author: Dixon Begay
func getLinksFromPage(URL string) string {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println("UTF-8? ", utf8.Valid(bytes))
	fmt.Println("HTML: ", string(bytes))
	parseHTMLforAnchorTags(resp.Body)
	return string(bytes)
}

// Authors: Timmoth Spate, Dixon Begay
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

// Author: Dixon Begay
func crawlPage(pageURL string) {
	if !stringInSlice(pageURL, crawled) {
		fmt.Println("Now crawling: ", pageURL)
		crawled = append(crawled, pageURL) // In the future this will be
		getLinksFromPage(pageURL)
	}
}

/*
Crawl crawls a website at the given url and will start
a number of concurrent processes depending on the threads input

Author: Dixon Begay, Timmothy Spate

Params:
	url <string> -----
		http address to website, can be IP or domain name, Ex. 71stsog.com, 12.345.45.67, http://google.com/

	threads <int> ----
		integer to specifiy the amount of concurrent processes to run
*/
func Crawl(rawURL string, threads int) {
	fmt.Println("go_crawl Version: ", version)
	url := getProperURL(rawURL)
	fmt.Println("URL: ", url)
	projectName = url.Host[:strings.Index(url.Host, ".")]
	fmt.Println("Project Name: ", projectName)
	createDirs()
	crawlPage(url.String())
	fmt.Println("Crawled slice:\n", crawled)
	// for i := 0; i <= threads; i++ {
	// 	go crawlPage(links[0])
	// }

	fmt.Println("Done...")
	os.Exit(0)
}
