package main

import (
	"fmt"
	"link-parser/link"
	"os"
)

func main() {
	f, err := os.Open("ex2.html")
	if err != nil {
		panic(err)
	}
	t, _ := link.Parse(f)
	fmt.Printf("%+v\n", t)
}
