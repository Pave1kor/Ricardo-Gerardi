package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}

func count(file io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(file)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
