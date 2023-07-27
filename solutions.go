package main

import (
	"fmt"
	"strings"
)

const stopOnFirstResult = true

func solutionWorker(solnFound <-chan []uint32, doneChan chan<- bool) {
	t := newTimer("solution parsing")

	defer func() {
		doneChan <- true
		t.end()
		return
	}()

	if stopOnFirstResult {
		soln, chanOpen := <-solnFound

		if !chanOpen {
			return
		}

		printSolution(soln)
	} else {
		for {
			soln, chanOpen := <-solnFound

			if !chanOpen {
				return
			}

			printSolution(soln)
		}
	}
}

func printSolution(soln []uint32) {
	fmt.Print("Solution found:")

	for _, binaryWord := range soln {
		fmt.Printf("\t%s", strings.Join(wordsList[binaryWord], "/"))
	}

	fmt.Printf("\n")
}
