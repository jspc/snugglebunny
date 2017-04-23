package main

import (
	"fmt"
	"os"
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

func (e editor) processKeyPress() int {
	var s string
	var k int = readKey()

	switch {
	case k == ctrlKey('s'):
		s = "snuggling to disk"
	case k == ctrlKey('q'):
		return exit
	case k == homeKey:
		s = "<home key>"
	case k == endKey:
		s = "<end key>"
	case k == backspace || k == delKey || k == ctrlKey('h'):
		s = "<delete>"
	case k == pageDown || k == pageUp:
		s = "<page up/ down>"
	case k == arrowUp || k == arrowDown || k == arrowLeft || k == arrowRight:
		s = "<cursor move>"
	case k == ctrlKey('l') || k == '\x1b':
		s = "<ignorable>"
	case k == '\r':
		s = "<new line>"
	default:
		s = string(k)
	}

	editorMessage = fmt.Sprintf("press: %d; evaluated to: %q", k, s)
	return cont
}

func readKey() int {
	var keybuf []byte = make([]byte, 1)
	_, _ = os.Stdin.Read(keybuf)

	if keybuf[0] == '\x1b' {
		var seq []byte = make([]byte, 3)
		_, _ = os.Stdin.Read(seq)

		if seq[0] == 0 || seq[1] == 0 {
			return int(keybuf[0])
		}

		if seq[0] == '[' {
			if seq[1] >= '0' && seq[1] <= '9' {
				if seq[2] == 0 {
					return '\x1b'
				}

				if seq[2] == '~' {
					switch seq[1] {
					case '1':
						return homeKey
					case '3':
						return delKey
					case '4':
						return endKey
					case '5':
						return pageUp
					case '6':
						return pageDown
					case '7':
						return homeKey
					case '8':
						return endKey
					}
				}

			} else {
				switch seq[1] {
				case 'A':
					return arrowUp
				case 'B':
					return arrowDown
				case 'C':
					return arrowRight
				case 'D':
					return arrowLeft
				case 'H':
					return homeKey
				case 'F':
					return endKey
				}
			}
		}

		if seq[0] == 'O' {
			switch seq[1] {
			case 'H':
				return homeKey
			case 'F':
				return endKey
			}
		}
	}
	return int(keybuf[0])
}

func ctrlKey(i int) int {
	return (i) & 0x1f
}
