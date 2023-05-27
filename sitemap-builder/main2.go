package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	lnk "sitemap/link"
)

var maxDepth int

type Link struct {
	Loc string `xml:"loc"`
}

type Urlset struct {
	Urls []Link `xml:"url"`
}

func main() {
	// Take url as cmd argument
	URL := flag.String("url", "", "Enter the root page url of the Domain.")
	MAXDEPTH := flag.Int("maxDepth", 1, "Max depth to traverse.")
	flag.Parse()

	maxDepth = *MAXDEPTH
	// Define a list and map to store the links
	links := []string{*URL}

	visited := map[string]struct{}{}

	// Fetch the home page and append the links in the list
	// Repeat the process for the links(not visited) the list

	bfs(*URL, &links, &visited, maxDepth)

	// Create XML object
	toXML := Urlset{
		Urls: make([]Link, len(links)),
	}
	for i, l := range links {
		toXML.Urls[i] = Link{l}
	}

	// Write in File

	data, err := xml.MarshalIndent(toXML, "", "	")
	ErrorHandler(err)
	ErrorHandler(os.WriteFile("test2.txt", data, 0777))
}

func bfs(link string, links *[]string, visited *map[string]struct{}, depth int) {

	// Do not fetch same page again
	if _, ok := (*visited)[link]; ok {
		return
	}
	pageLinks := get(link)

	// add link to visited map and append found links to the list
	(*visited)[link] = struct{}{}
	*links = append(*links, pageLinks...)

	// Call bfs for each found link if depth is > 1
	if depth > 1 {
		for _, l := range pageLinks {
			bfs(l, links, visited, depth-1)
		}
	}
}

func get(pageLink string) []string {

	found := []string{}
	resp, err := http.Get(pageLink)
	ErrorHandler(err)

	herfTexts, err := lnk.Parse(resp.Body)
	ErrorHandler(err)
	base := (&url.URL{
		Scheme: resp.Request.URL.Scheme,
		Host:   resp.Request.URL.Host,
	}).String()

	// Append all the link to the list
	for _, ht := range herfTexts {

		link := ht.Href
		switch {
		case strings.HasPrefix(link, "/"):
			link = base + link
		case strings.HasPrefix(link, "http"):
			link = link
		default:
			continue
		}
		found = append(found, link)
	}

	return found
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println("Something went wrong", err)
		os.Exit(1)
	}
}
