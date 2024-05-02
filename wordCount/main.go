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
	result := count(os.Stdin, *lines)
	fmt.Println("hello")
	fmt.Println(result)
}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	} else {
		scanner.Split(bufio.ScanLines)
	}

	res := 0
	for scanner.Scan() {
		res++
	}
	return res
}
