package gogit

import (
	"bytes"
	"fmt"
)

type ContentType string

const (
	ContentType_Blob   ContentType = "blob"
	ContentType_Commit ContentType = "commit"
	ContentType_Tree   ContentType = "tree"
)

func Header(contentType ContentType, length int) ([]byte, error) {
	var b bytes.Buffer
	_, err := fmt.Fprintf(&b, "%s %d", contentType, length)
	if err != nil {
		return nil, err
	}
	if err := b.WriteByte(0); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
