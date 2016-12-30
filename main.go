package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Global Variables
var base_url string
var project_name string
var queue_file string
var crawled_file string
var errors_file string
var summary_file string
var num_pages int
var num_errors int
var num_pdf int
var num_html int
var num_media int
var num_other int

func main() {
	url := input()
	Domain_name := get_domain_name(url)
	fmt.Println("Domain Name: ", Domain_name)
	project_name = Domain_name[:strings.Index(Domain_name, ".")]
	fmt.Println(project_name)

	fmt.Println("The path is !!!!XXXX:   " + filepath.Dir("."+string(filepath.Separator)+"projects"+string(filepath.Separator)+project_name))
	create_dirs()
}
