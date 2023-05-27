package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	file := flag.String("csv", "problems.csv", "A scv file in format of 'question,answer'.")
	limit := flag.Int("limit", 20, "The time limit for the quiz in seconds.")
	flag.Parse()

	score := 0
	f, err := os.Open(*file)
	if err != nil {
		log.Fatal("Failed to open file:", *file)
	}
	r := csv.NewReader(f)

	if err != nil {
		log.Fatal("Failed to parse csv file. ", err)
	}
	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	qa, err := r.Read()
	for i := 1; err == nil; i++ {
		fmt.Printf("Que %v: %v = ?\nYour ans: ", i, qa[0])
		ip_chan := make(chan string)
		go func() {
			var guess string
			fmt.Scanf("%s\n", &guess)
			ip_chan <- guess
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTIME OUT.")
			return
		case res := <-ip_chan:
			if strings.TrimSpace(res) == strings.TrimSpace(qa[1]) {
				score++
				fmt.Printf("Correct!\n\n")
			}
		}

		qa, err = r.Read()
	}

	fmt.Println("Your total score:", score)

}
