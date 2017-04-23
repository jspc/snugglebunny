package main

import (
	"os"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

type window struct {
	title []byte

	rows      int
	cols      int
	cursorRow int
	cursorCol int
}

func NewWindow() window {
	cols, rows, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	return window{
		rows:  rows,
		cols:  cols,
		title: []byte("snugglebunny <3"),
	}
}

func (w window) editorPane(contents []string, start int) (s []string) {
	s = make([]string, w.rows-3)

	for i := 0; i < w.rows-3; i++ {
		if (i + start) < len(contents) {
			s[i] = contents[i+start]
		} else {
			s[i] = ""
		}
	}
	return
}

func (w window) titleBar() string {
	b := make([]byte, w.cols)

	bar := color.New(color.Bold).Add(color.FgHiMagenta).Add(color.BgHiWhite)

	startPoint := (w.cols / 2) - (len(w.title) / 2)

	for idx, c := range w.title {
		b[startPoint+idx] = c
	}

	b = cleanStringNil(b)

	return bar.Sprintf(string(b))
}

func (w window) fileBar(filename, suffix string, dirty bool, bPos, bLen int) string {
	b := make([]byte, w.cols)

	bar := color.New(color.Reset).Add(color.FgHiBlack).Add(color.BgHiWhite)

	for idx, c := range []byte(filename) {
		b[idx] = c
	}

	if dirty {
		for idx, c := range dirtyMessage {
			b[idx+len(dirtyMessage)+1] = c
		}
	}

	suffixStatus := []byte("[" + suffix + "]")
	for idx, c := range suffixStatus {
		b[w.cols-len(suffixStatus)+idx] = c
	}

	b = cleanStringNil(b)

	return bar.Sprintf(string(b))

}

func (w window) messageBar() string {
	bar := color.New(color.Reset).Add(color.FgHiCyan)

	return bar.Sprintf(editorMessage)
}

func cleanStringNil(b []byte) []byte {
	newB := make([]byte, len(b))

	for idx, c := range b {
		if c == 0 {
			newB[idx] = 32
		} else {
			newB[idx] = c
		}
	}

	return newB
}
