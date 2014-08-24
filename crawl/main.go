package main

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode/utf8"
)

const alphabet string = "абвгдежзийклмнопрстуфхцчшщъыьэюяё-"
const alphabetVowels string = "аеиоуыэюяё"

var alphabetLength int

func randomWords() []string {
	random := rand.New(rand.NewSource(1234))
	randomChoice := func() rune {
		pos := random.Intn(alphabetLength)
		i := 0
		for _, ch := range alphabet {
			if i == pos {
				return ch
			}
			i++
		}
		panic("unexpected pos out of range")
	}
	randomWord := func(length int) string {
		s := make([]rune, 0, length+1)
		var ch rune
		hasAccent := false
		for i := 0; i < length; i++ {
			ch = randomChoice()
			s = append(s, ch)
			if !hasAccent && strings.IndexRune(alphabetVowels, ch) != -1 && random.Intn(10) == 0 {
				s = append(s, '*')
				hasAccent = true
			}
		}
		return string(s)
	}

	result := make([]string, 0, 400)

	// singular letters
	for _, ch := range alphabet {
		result = append(result, string(ch))
	}

	// 2 letter words
	for i := 1; i <= 100; i++ {
		result = append(result, randomWord(2))
	}

	// 3 letter words
	for i := 1; i <= 100; i++ {
		result = append(result, randomWord(3))
	}

	// 4 letter words
	for i := 1; i <= 100; i++ {
		result = append(result, randomWord(4))
	}

	return result
}

func main() {
	alphabetLength = utf8.RuneCountInString(alphabet)

	words := randomWords()
	fmt.Printf("words: %v\n", words)

	for _, w := range words {
		item, err := CrawlWord(w)
		check(err)

		versions, _ := items[w]
		versions = append(versions, *item)
		items[w] = versions

		storageSave()
	}
}
