package main

import (
    "bufio"
    "os"

    _"golang.org/x/crypto/ssh/terminal"
)

func main() {
    b := NewBuffer("~/go/src/github.com/jspc/snugglebunny/buffer.go")
    e := NewEditor()
    e.buffers = append(e.buffers, b)

    // oldState, err := terminal.MakeRaw(0)
    // if err != nil {
    //     panic(err)
    // }
    // defer terminal.Restore(0, oldState)

    e.draw()
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}
