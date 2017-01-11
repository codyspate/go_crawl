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

func appendToFile(path string, data string) {
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
func deleteFileContents(path string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()
}

// Create queue and crawled files (if not created)
func createDataFiles(projectName string, base_url string) {
	queue := filepath.Join(".\\projects", projectName, "queue.txt")
	crawled := filepath.Join(".\\projects", projectName, "crawled.txt")
	summary := filepath.Join(".\\projects", projectName, "summary.txt")
	errors := filepath.Join(".\\projects", projectName, "errors.txt")
	fmt.Println(queue)
	fmt.Println(crawled)
	fmt.Println(summary)
	fmt.Println(errors)
	deleteFileContents(queue)
	deleteFileContents(crawled)
	deleteFileContents(summary)
	deleteFileContents(errors)
	if exists(queue) {
		writeFile(queue, base_url)
	}
	if exists(crawled) {
		writeFile(crawled, "")
	}
	if exists(summary) {
		t := "PDF count:     \nHTML/HTM count:     \nMedia files:     \nOther:     \nErrors:     \n\nTotal size:     "
		writeFile(summary, t)
	}
	if exists(errors) {
		writeFile(errors, "")
	}
}

func updateSummary(
	path string, pname string, url string,
	pdf string, html string, media string,
	other string, errors string, pages string,
	size string, queue string, crawled string) {
	deleteFileContents(path)
	appendToFile(path, fmt.Sprintf("Cal Crawler v%s\n\n", VERSION))
	appendToFile(path, fmt.Sprintf("Website: %s\n", pname))
	appendToFile(path, fmt.Sprintf("URL: %s\n\n", url))
	appendToFile(path, fmt.Sprintf("PDF count: %s\nHTML/HTML count: %s\nMedia files: %s\nOther: %s\nErrors: %s\n\nTotal Number of Pages: %s\nTotal size: %sMB\n\n", pdf, html, media, other, errors, pages, size))
	appendToFile(path, fmt.Sprintf("Queue: %s\nCrawled: %s\n\n", queue, crawled))
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
