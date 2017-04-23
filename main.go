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
	b := NewBuffer("/Users/jspc/golang/src/github.com/jspc/snugglebunny/buffer.go")
	e := NewEditor()

	err := e.loadBuffer(b)
	if err != nil {
		panic(err)
	}

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
	print("\r\n")
}
