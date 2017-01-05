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

func write_file(path string, text string) {
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

func append_to_file(path string, data string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(data); err != nil {
		panic(err)
	}
}

func create_dirs() {
	var path bytes.Buffer
	path.WriteString("projects" + string(filepath.Separator) + project_name)
	fmt.Println("The path is !!!!XXXX: " + path.String())
	os.MkdirAll(path.String(), os.ModePerm)
	// check(e)
	queue_file = filepath.Join(path.String(), "queue.txt")
	crawled_file = filepath.Join(path.String(), "crawled.txt")
	errors_file = filepath.Join(path.String(), "errors.txt")
	summary_file = filepath.Join(path.String(), "summary.txt")
	a := [5]string{queue_file, crawled_file, errors_file, summary_file}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		f, _ := os.Create(a[i])
		// check(e)
		f.Close()
	}
}

// Delete the contents of a file
func delete_file_contents(path string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()
}

// Create queue and crawled files (if not created)
func create_data_files(project_name string, base_url string) {
	queue := filepath.Join(".\\projects", project_name, "queue.txt")
	crawled := filepath.Join(".\\projects", project_name, "crawled.txt")
	summary := filepath.Join(".\\projects", project_name, "summary.txt")
	errors := filepath.Join(".\\projects", project_name, "errors.txt")
	fmt.Println(queue)
	fmt.Println(crawled)
	fmt.Println(summary)
	fmt.Println(errors)
	delete_file_contents(queue)
	delete_file_contents(crawled)
	delete_file_contents(summary)
	delete_file_contents(errors)
	if exists(queue) {
		write_file(queue, base_url)
	}
	if exists(crawled) {
		write_file(crawled, "")
	}
	if exists(summary) {
		t := "PDF count:     \nHTML/HTM count:     \nMedia files:     \nOther:     \nErrors:     \n\nTotal size:     "
		write_file(summary, t)
	}
	if exists(errors) {
		write_file(errors, "")
	}
}
