package main

import (
	"golang.org/x/crypto/ssh/terminal"
)

var (
	oldState *terminal.State
	err      error
)

const (
	exit = iota
	cont
)

func main() {
	b := NewBuffer("~/go/src/github.com/jspc/snugglebunny/buffer.go")
	e := NewEditor()
	e.buffers = append(e.buffers, b)

	rawMode()
	defer endRawMode()

	for {
		e.draw()
		if e.processKeyPress() == exit {
			return
		}
	}
}

func rawMode() {
	oldState, err = terminal.MakeRaw(1)
	if err != nil {
		panic(err)
	}
}

func endRawMode() {
	terminal.Restore(1, oldState)
}
