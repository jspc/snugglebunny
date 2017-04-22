package main

import (
    "path"
)

type buffer struct{
    filename string
    filepath string
    suffix string
    dirty bool

    contents []string
}

func NewBuffer(filepath string) buffer {
    return buffer{
        filename: path.Base(filepath),
        filepath: filepath,
        suffix: path.Ext(filepath),
    }
}
