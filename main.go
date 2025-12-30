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
	c := Count(os.Stdin)
	fmt.Printf("%8d %7d %7d\n", c.Lines, c.Words, c.Bytes)
}
