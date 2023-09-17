package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	if len(os.Args) <= 1 {
		panic(fmt.Sprintf("Usage: %s <filename>", os.Args[0]))
	}

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r, err := zlib.NewReader(f)
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	spew.Dump(content)
}
