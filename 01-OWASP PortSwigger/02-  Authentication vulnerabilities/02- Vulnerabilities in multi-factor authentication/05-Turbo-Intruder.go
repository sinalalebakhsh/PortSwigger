package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/PortSwigger/turbo-intruder/pkg/burp"
)

func queueRequests(target *burp.Target, wordlists []string) {
	engine := burp.NewRequestEngine(target.Endpoint, 5, 100, false)

	for i := 3; i < 8; i++ {
		engine.Queue(target.Req, randstr(i), 1)
		engine.Queue(target.Req, target.BaseInput, 2)
	}

	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		engine.Queue(target.Req, word, 0)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}
}

func handleResponse(req *burp.Request, interesting bool) {
	if interesting {
		table.Add(req)
	}
}

func main() {
	target := &burp.Target{} // Replace with your actual target

	// Populate wordlists if necessary
	wordlists := []string{}

	queueRequests(target, wordlists)
}