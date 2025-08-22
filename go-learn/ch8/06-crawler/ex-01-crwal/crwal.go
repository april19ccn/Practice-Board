// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	links "example/learn/ch5/06-anonymous/03-links"
	"fmt"
	"log"
	"os"
)

type EXData struct {
	depth int
	data  []string
}

// !+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string, depth int) EXData {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return EXData{depth + 1, list}
}

//!-sema

// !+
func main() {
	worklist := make(chan EXData)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- EXData{0, os.Args[1:]} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		if t := <-worklist; t.depth < 2 {
			for _, link := range t.data {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string, depth int) {
						worklist <- crawl(link, depth)
					}(link, t.depth)
				}
			}
		} else {
			fmt.Println("大于2访问层级了!!")
		}
	}
}

//!-
