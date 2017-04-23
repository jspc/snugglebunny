package main

import (
	"fmt"
)

const (
	backspace = 127
	arrowLeft = 1000 + iota
	arrowRight
	arrowUp
	arrowDown
	delKey
	homeKey
	endKey
	pageUp
	pageDown
)

func processKeyPress(b []byte) {
	editorMessage = fmt.Sprintf("%d", b[0])
}
