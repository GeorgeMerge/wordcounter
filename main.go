package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"unicode"
)

type (
	Counts struct {
		Name  string `json:"name,omitempty"`
		Lines int    `json:"lines"`
		Words int    `json:"words"`
		Bytes int    `json:"bytes"`
	}

	IndexedResult struct {
		Index  int
		Counts Counts
	}
)

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

func countMultiple(filenames []string) []Counts {
	results := make([]Counts, len(filenames))

	var wg sync.WaitGroup
	resultCh := make(chan IndexedResult, len(filenames))

	for i, filename := range filenames {
		wg.Add(1)
		go func(idx int, fname string) {
			defer wg.Done()

			file, err := os.Open(fname)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			c := Count(file)
			c.Name = fname
			resultCh <- IndexedResult{Index: idx, Counts: c}
		}(i, filename)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		results[res.Index] = res.Counts
	}

	return results
}

func printJSON(results []Counts) {
	output, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}

func main() {
	if len(os.Args) > 1 {
		filenames := os.Args[1:]
		printJSON(countMultiple(filenames))
	} else {
		c := Count(os.Stdin)
		printJSON([]Counts{c})
	}
}
