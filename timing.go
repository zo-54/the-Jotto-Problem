package main

import (
	"fmt"
	"strings"
	"time"
)

type timer struct {
	label     string
	startTime time.Time
}

func newTimer(label string) *timer {
	t := &timer{
		label:     strings.ToUpper(label),
		startTime: time.Now(),
	}

	return t
}

func (t timer) end() {
	fmt.Printf("Completed: %s in %dms\n", t.label, time.Since(t.startTime).Milliseconds())
}
