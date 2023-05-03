package main

// All lowercase english letters in order of how common they are (descending)
const commonLetters = "aesiorunltycdhmpgkbwfvzjxq"

var (
	letterMap map[rune]uint32
	wordsList map[uint32][]string
)

func main() {
	programTimer := newTimer("program")

	// Set up a map of runes (letters) to their binary representation
	letterMap = make(map[rune]uint32)

	for i, char := range commonLetters {
		letterMap[char] = 1 << i
	}

	// initialise channels
	wordChan := make(chan uint32)
	solnChan := make(chan []uint32)
	doneChan := make(chan bool)

	// start workers
	go fileReader(wordChan)
	go treeWorker(wordChan, solnChan)
	go solutionWorker(solnChan, doneChan)

	<-doneChan

	programTimer.end()
}
