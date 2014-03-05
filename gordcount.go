package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var filename = flag.String("file", "", "filename to read from")

func WordCount(r io.Reader) map[string]int64 {

	scanner := bufio.NewScanner(r)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	counts := make(map[string]int64)

	for scanner.Scan() {
		word := scanner.Text()
		_, ok := counts[word]

		if ok == true {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	return counts
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: <exec> filename thresholdcount")
	}

	threshold, err := strconv.ParseInt(os.Args[2], 10, 8)
	if err != nil {
		log.Fatal("Usage: <exec> filename thresholdcount")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	counts := WordCount(file)

	for key, _ := range counts {
		if counts[key] > threshold {
			fmt.Printf("%7d: %s\n", counts[key], key)
		}
	}
}
