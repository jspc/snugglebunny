package main

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	oldState *terminal.State
	err      error
)

func main() {
	b := NewBuffer("~/go/src/github.com/jspc/snugglebunny/buffer.go")
	e := NewEditor()
	e.buffers = append(e.buffers, b)

	rawMode()
	defer endRawMode()

	var keybuf []byte = make([]byte, 1)
	for {
		e.draw()
		_, _ = os.Stdin.Read(keybuf)
		processKeyPress(keybuf)
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
