package go_crawl

import (
	"fmt"
	"os"
	"path/filepath"
	//"io/ioutil"
)

// func input() string {
// 	scan := bufio.NewScanner(os.Stdin)
// 	fmt.Print("Input URL: ")
// 	scan.Scan()
// 	Homepage := scan.Text()
// 	if !(strings.Contains(Homepage, "://")) {
// 		Homepage = "http://" + Homepage
// 		if !(strings.Contains(string(Homepage[len(Homepage)-1]), "/")) {
// 			Homepage = Homepage + "/"
// 		}
// 	}
// 	return Homepage
// }

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

func update_summary(
	path string, pname string, url string,
	pdf string, html string, media string,
	other string, errors string, pages string,
	size string, queue string, crawled string) {
	delete_file_contents(path)
	append_to_file(path, fmt.Sprintf("Cal Crawler v%s\n\n", VERSION))
	append_to_file(path, fmt.Sprintf("Website: %s\n", pname))
	append_to_file(path, fmt.Sprintf("URL: %s\n\n", url))
	append_to_file(path, fmt.Sprintf("PDF count: %s\nHTML/HTML count: %s\nMedia files: %s\nOther: %s\nErrors: %s\n\nTotal Number of Pages: %s\nTotal size: %sMB\n\n", pdf, html, media, other, errors, pages, size))
	append_to_file(path, fmt.Sprintf("Queue: %s\nCrawled: %s\n\n", queue, crawled))
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Iterate through a set, each item will be a line in a file
// func set_to_file(links string, file_name string) {

//     try:
//         with open(file_name,"w") as f:
//             for l in sorted(links):
//                 f.write(l+"\n")
//     except Exception as e:
//         print(str(e))
// }
