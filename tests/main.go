package main

import (
	"flag"
	"os"

	"github.com/codyspate/go_crawl"
)

func main() {
	urlPtr := flag.String("url", "http://www.google.com/", "a url <string>")
	threadPtr := flag.Int("threads", 10, "number of threads <int>")
	flag.Parse()
	// Below is from Python version
	// 71stsog | http://71stsog.com/ | 71stsog.com
	// Spider(PROJECT_NAME, HOMEPAGE, DOMAIN_NAME)
	go_crawl.Crawl(*urlPtr, *threadPtr)
	os.Exit(0)
}
