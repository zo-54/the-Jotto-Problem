package main

import (
	"fmt"
	"sort"
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

		plainTextSolution := make([]string, 5, 5)

		for i := range soln {
			plainTextSolution[i] = wordsList[soln[i]][0]
		}

		sort.Strings(plainTextSolution)

		fmt.Println(strings.Join(plainTextSolution, ", "))
	} else {
		// TODO: Implement
		panic("not implemented!")

		for {
			_, chanOpen := <-solnFound

			if !chanOpen {
				return
			}
		}
	}
}
