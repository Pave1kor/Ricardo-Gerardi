package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type calc struct {
	result map[string]int
}

type Counter interface {
	countSymbol(*bufio.Scanner, string)
}

// Реализация метода интерфейса для calc
func (c *calc) countSymbol(scanner *bufio.Scanner, symbol string) {
	count := 0
	for scanner.Scan() {
		count++
	}
	if c.result == nil {
		c.result = make(map[string]int)
	}
	c.result[symbol] = count
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(file io.Reader, countLines, countBytes bool) int {
	scanner := bufio.NewScanner(file)

	// Создаем экземпляр структуры, реализующей интерфейс
	counter := &calc{result: make(map[string]int)}
	switch {
	case countLines:
		scanner.Split(bufio.ScanLines)
		counter.countSymbol(scanner, "line")
		return counter.result["line"]
	case countBytes:
		scanner.Split(bufio.ScanBytes)
		counter.countSymbol(scanner, "bytes")
		return counter.result["bytes"]
	default:
		scanner.Split(bufio.ScanWords)
		counter.countSymbol(scanner, "words")
		return counter.result["words"]
	}
}

//можно обавить интерфейсы и методы, исключить или разрешить использование нескольких флагов.
