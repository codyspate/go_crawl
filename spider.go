package main

type Spider struct {
	project_name, base_url, domain_name, queue_file, crawled_file, summary_file, errors_file string
	num_pdf, num_html, num_media, num_other, num_errors, total_size, pages                   int
}
