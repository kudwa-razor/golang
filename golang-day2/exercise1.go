package main

import (
	"fmt"
	"sync"
)

func CounLetters(word string, freqMap chan map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()
	freq := make(map[rune]int)
	for _, char := range word {
		freq[char]++
	}
	freqMap <- freq
}

func main() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	freqMap := make(chan map[rune]int, len(words))
	var wg sync.WaitGroup

	for _, word := range words {
		wg.Add(1)
		go CounLetters(word, freqMap, &wg)
	}

	go func() {
		wg.Wait()
		close(freqMap)
	}()

	finalfreq := make(map[rune]int)
	for freq := range freqMap {
		for char, count := range freq {
			finalfreq[char] += count
		}
	}
	for char, count := range finalfreq {
		fmt.Printf("%c: %d\n", char, count)
	}
}
