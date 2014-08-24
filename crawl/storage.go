package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/temoto/phonetic-semantics/phonosem"
)

const dbPath = "phonosem.data"

type Item struct {
	phonosem.Item
	ModifiedAt time.Time
}

var (
	items map[string][]Item
)

func check(err error) {
	if err != nil {
		// _, file, line, _ := runtime.Caller(1)
		panic(err)
	}
}

func storageRead() error {
	panic("not implemented")
}

func storageSave() error {
	f, err := os.Create("phonosem-crawl-db-v1")
	check(err)
	defer f.Close()

	fb := bufio.NewWriter(f)
	defer fb.Flush()
	for word, is := range items {
		for _, item := range is {
			fmt.Fprintf(fb, `"%s","%s"`, word, item.ModifiedAt.String())
			for _, f := range item.Features {
				fmt.Fprintf(fb, ",%d,%d", f.Id, f.Value)
			}
			fb.WriteString("\n")
		}
	}

	return nil
}

func init() {
	items = make(map[string][]Item)
	// check(storageRead())
}
