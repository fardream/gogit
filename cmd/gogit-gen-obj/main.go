package main

import (
	"flag"
	"os"
	"path"

	"github.com/fardream/gogit"
)

func main() {
	var content, prefix string

	flag.StringVar(&content, "content", "new content", "content of the file")
	flag.StringVar(&prefix, "prefix", "", "path to git repo")

	flag.Parse()

	blob, err := gogit.NewBlob([]byte(content))
	if err != nil {
		panic(err)
	}

	hexDigest := blob.HexDigest()
	dirName := hexDigest[0:2]
	fileName := hexDigest[2:]

	fullDir := path.Join(prefix, ".git", "objects", dirName)
	if err := os.MkdirAll(fullDir, 0o777); err != nil {
		panic(err)
	}

	fullFilePath := path.Join(fullDir, fileName)
	if err := os.WriteFile(fullFilePath, blob.Blob, 0o666); err != nil {
		panic(err)
	}
}
