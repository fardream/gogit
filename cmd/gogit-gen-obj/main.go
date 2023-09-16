package main

import (
	"os"
	"path"

	"github.com/fardream/gogit"
)

func main() {
	content := "new file\n"

	blob, err := gogit.NewBlob([]byte(content))
	if err != nil {
		panic(err)
	}

	dirName := blob.HexDigest[0:2]
	fileName := blob.HexDigest[2:]

	prefix := "/Users/fardream/github/test/test"

	fullDir := path.Join(prefix, ".git", "objects", dirName)
	if err := os.MkdirAll(fullDir, 0o777); err != nil {
		panic(err)
	}

	fullFilePath := path.Join(fullDir, fileName)
	if err := os.WriteFile(fullFilePath, blob.Blob, 0o666); err != nil {
		panic(err)
	}
}
