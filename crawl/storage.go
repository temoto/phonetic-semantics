package main

import (
	"github.com/temoto/phonetic-semantics/phonosem"
)

const dbPath = "phonosem.data"

var (
	items map[string][]phonosem.Item
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
	panic("not implemented")
}

func init() {
	items = make(map[string][]phonosem.Item)
	// check(storageRead())
}
