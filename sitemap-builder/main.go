package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	lnk "sitemap/link"
)

type Link struct {
	Loc string `xml:"loc"`
}

type Urlset struct {
	Urls []Link `xml:"url"`
}

func main() {
	// Take url as cmd argument
	URL := flag.String("url", "", "Enter the root page url of the Domain.")
	maxDepth := flag.Int("maxDepth", 1, "Max depth to traverse.")
	flag.Parse()

	// Define a list and map to store the links
	links := []string{*URL}

	linkMap := map[string]int{*URL: 0}

	// Fetch the home page and append the links in the list
	// Repeat the process for the links(not visited) the list
	for i := 0; i < len(links) && linkMap[links[i]] < *maxDepth; i++ {
		resp, err := http.Get(links[i])
		ErrorHandler(err)

		level := linkMap[links[i]]

		herfTexts, err := lnk.Parse(resp.Body)
		ErrorHandler(err)
		base := (&url.URL{
			Scheme: resp.Request.URL.Scheme,
			Host:   resp.Request.URL.Host,
		}).String()

		// Append all the link to the list
		for _, ht := range herfTexts {

			link := ht.Href
			if link[0] == '#' {
				continue
			}
			if link[0] == '/' {
				link = base + link
			}
			if _, ok := linkMap[ht.Href]; !ok {
				links = append(links, link)
				linkMap[link] = level + 1
			}
		}
	}

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
	ErrorHandler(os.WriteFile("test.txt", data, 0777))
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println("Something went wrong", err)
		os.Exit(1)
	}
}
