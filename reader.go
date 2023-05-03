package main

import (
	"bufio"
	"log"
	"math/bits"
	"os"
	"path/filepath"
)

// TODO: Read all words first so anagrams aren't lost
func fileReader(addWord chan<- uint32) {
	t := newTimer("reading file")

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

		addWord <- binaryWord
	}

	close(addWord)

	t.end()
}
