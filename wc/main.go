package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(reader io.Reader, linesCount bool, byteCount bool) (int, error) {
	wc := 0
	scanner := bufio.NewScanner(reader)
	if linesCount {
		scanner.Split(bufio.ScanLines)
	} else if byteCount {
		fmt.Println("SPLIT by bytes")
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanWords)
	}
	for scanner.Scan() {
		wc++
	}
	return wc, nil
}

func main() {
	lineCount := flag.Bool("l", false, "Count lines")
	byteCount := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	wc, err := count(os.Stdin, *lineCount, *byteCount)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var postFix string
	switch {
	case *lineCount:
		postFix = "lines"
	case *byteCount:
		postFix = "bytes"
	default:
		postFix = "words"
	}

	fmt.Println(wc, postFix)
}
