package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"path/filepath"
	"sort"
)

func fileReader(addWord chan<- uint32) {
	t := newTimer("reading file")

	var words []uint32

	wordsList = make(map[uint32][]string)

	filePath, err := filepath.Abs("./words_alpha.txt")
	if err != nil {
		log.Fatalf("could not parse file path")
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not parse file path")
	}

	// close file on function completion
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 5 {
			continue
		}

		var binaryWord uint32 = 0

		for _, char := range line {
			binaryWord = binaryWord | letterMap[char]
		}

		if bits.OnesCount32(binaryWord) != 5 {
			continue
		}

		wordsList[binaryWord] = append(wordsList[binaryWord], line)
	}

	for k := range wordsList {
		words = append(words, k)
	}

	sort.Slice(words, func(i, j int) bool { return words[i] > words[j] })

	t.end()

	fmt.Printf("Found %v words (not including anagrams).\n", len(words))

	for _, w := range words {
		addWord <- w
	}

	close(addWord)
}
