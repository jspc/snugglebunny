package main

import (
	"os"

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
	e := NewEditor()

	if len(os.Args) > 1 {
		for _, p := range os.Args[1:] {
			b := NewBuffer(p)

			err := e.loadBuffer(b)
			if err != nil {
				panic(err)
			}
		}
	} else {
		e.loadBuffer(&buffer{})
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
