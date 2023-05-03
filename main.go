package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const lettersByMostCommon = "aesiorunltycdhmpgkbwfvzjxq"

// Create map of runes to their corresponding binary representation
var (
	letterMap map[rune]uint32
	groups    []*group
	wordsList map[uint32][]string
	solutions [][]uint32
)

type group struct {
	c uint32
	l []uint32
}

func main() {
	// Begin timing
	startTiming()

	letterMap = make(map[rune]uint32)

	// Set up a map of runes (letters) to their binary representation
	for i, char := range lettersByMostCommon {
		letterMap[char] = 1 << i
	}

	// Create wait group
	wg := &sync.WaitGroup{}

	// Set up word channels
	addWord := make(chan uint32)

	wg.Add(1)
	go fileReader(wg, addWord)

	wg.Add(1)
	go wordsWorker(wg, addWord)

	wg.Wait()

	solutions = make([][]uint32, 0)

	for _, g := range groups {
		if bits.OnesCount32(g.c) > 24 {
			solutions = append(solutions, g.l)
		}
	}

	endTiming()

	for _, s := range solutions {
		fmt.Println(s)
	}

	fmt.Println(len(wordsList))
	fmt.Println(len(groups))
	fmt.Println(len(solutions))
}

func fileReader(wg *sync.WaitGroup, addWord chan<- uint32) {
	defer func() {
		wg.Done()
		close(addWord)
		checkpoint("reading file")
	}()

	wordsList = make(map[uint32][]string)

	filePath, err := filepath.Abs("./wordle_data/answers.txt")
	handleError("could not parse file path", err)

	file, err := os.Open(filePath)
	handleError("could not open file", err)
	defer CloseQuietly(file)

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
}

func wordsWorker(wg *sync.WaitGroup, addWord <-chan uint32) {
	defer func() {
		wg.Done()

		checkpoint("creating groups")
	}()

	groups = make([]*group, 1)

	groups[0] = &group{
		c: 0,
		l: make([]uint32, 5),
	}

	for {
		w, open := <-addWord

		for _, g := range groups {
			if g.c&w == 0 {
				list := make([]uint32, 5)

				for i := range g.l {
					list[i] = g.l[i]
				}

				groups = append(groups, &group{
					c: g.c | w,
					l: list,
				})
			}
		}

		if !open {
			break
		}
	}
}

func CloseQuietly(closer io.Closer) {
	_ = closer.Close()
}

func handleError(msg string, err error) {
	if err != nil {
		log.Fatalf(msg+", err: %v", err)
	}
}

// Timing
var (
	startTime time.Time
)

func startTiming() {
	startTime = time.Now()
}

func checkpoint(name string) {
	fmt.Printf("Completed %v in %dms\n", name, time.Since(startTime).Milliseconds())
}

func endTiming() {
	endTime := time.Since(startTime)

	fmt.Printf("Finished operation in %dms\n", endTime.Milliseconds())
}
