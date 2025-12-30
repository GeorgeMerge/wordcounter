package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type Counts struct {
	Lines int
	Words int
	Bytes int
}

func Count(r io.Reader) Counts {
	var c Counts
	inWord := false

	reader := bufio.NewReader(r)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}

		c.Bytes++

		if b == '\n' {
			c.Lines++
		}

		if unicode.IsSpace(rune(b)) {
			inWord = false
		} else if !inWord {
			inWord = true
			c.Words++
		}
	}

	return c
}

func main() {
	var input io.Reader = os.Stdin
	var filename string

	if len(os.Args) > 1 {
		filename = os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		input = file
	}

	c := Count(input)

	if filename != "" {
		fmt.Printf("%8d %7d %7d %s\n", c.Lines, c.Words, c.Bytes, filename)
	} else {
		fmt.Printf("%8d %7d %7d\n", c.Lines, c.Words, c.Bytes)
	}
}
