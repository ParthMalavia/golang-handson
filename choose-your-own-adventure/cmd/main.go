package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	st "example/cyoa/story"
)

func main() {
	PORT := flag.String("port", "3000", "Port number on which server will be running.")
	fileName := flag.String("file", "story.json", "Name of json file that contain Choose Your Own Adventure story.")
	flag.Parse()
	fmt.Println("Using story form %s file", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	story, err := st.JsonToStory(f)
	if err != nil {
		panic(err)
	}

	// Use diff template, pathFunc
	handler := st.GetNewHandler(story)
	fmt.Println("Running on port:", *PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *PORT), handler))
}
