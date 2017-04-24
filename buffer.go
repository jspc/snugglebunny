package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path"
)

type buffer struct {
	filename string
	filepath string
	suffix   string
	dirty    bool

	contents         []string
	contentsPosition int
}

func NewBuffer(filepath string) *buffer {
	return &buffer{
		filename: path.Base(filepath),
		filepath: filepath,
		suffix:   path.Ext(filepath),
	}
}

func (b *buffer) load() (err error) {
	if b.filepath == "" {
		return
	}

	var f *os.File
	var part []byte
	var prefix bool

	if f, err = os.Open(b.filepath); err != nil {
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			contents := b.contents
			contents = append(contents, buffer.String())

			b.contents = contents
			buffer.Reset()
		}
	}

	if err == io.EOF {
		err = nil
	}
	return
}
