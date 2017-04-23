package main

import (
	"fmt"
)

var (
	editorMessage = "i luv u"
	dirtyMessage  = []byte("[un-snuggled]")
)

type editor struct {
	window  window
	buffers []*buffer
	bIndex  int
}

func NewEditor() editor {
	return editor{
		window:  NewWindow(),
		buffers: []*buffer{},
	}
}

func (e *editor) loadBuffer(b *buffer) (err error) {
	err = b.load()
	if err != nil {
		return
	}

	e.buffers = append(e.buffers, b)
	e.bIndex = len(e.buffers) - 1

	return
}

func (e editor) draw() {
	buffer := e.buffers[e.bIndex]

	print("\033[H\033[2J")

	lines := make([]string, e.window.rows)
	lines[0] = e.window.titleBar()

	for idx, line := range e.window.editorPane(buffer.contents, buffer.contentsPosition) {
		lines[idx+1] = line
	}

	lines[len(lines)-2] = e.window.fileBar(buffer.filename, buffer.suffix, buffer.dirty, 0, len(buffer.contents))
	lines[len(lines)-1] = e.window.messageBar()

	for idx, l := range lines {
		fmt.Print(l)
		if idx < (e.window.rows - 2) {
			fmt.Print("\n\r")
		}
	}
}
