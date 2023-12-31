package gogit

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
)

type Object struct {
	// header, "blob length-of-content" and terminating null byte.
	Header []byte
	// digest is the 20 byte sha-1 hash of header + content
	Digest Sha1
	// Blob contains the actually zlib-compressed header + content of the blob.
	Blob []byte
}

func (o *Object) HexDigest() string {
	return o.Digest.String()
}

func NewObject(
	contentType ContentType,
	content []byte,
) (*Object, error) {
	l := len(content)
	header, err := NewHeader(contentType, l)
	if err != nil {
		return nil, err
	}
	hasher := sha1.New()
	_, err = hasher.Write(header)
	if err != nil {
		return nil, err
	}
	_, err = hasher.Write(content)
	if err != nil {
		return nil, err
	}
	digest := Sha1(hasher.Sum(nil)[:sha1Length])

	var b bytes.Buffer
	z := zlib.NewWriter(&b)
	_, err = z.Write(header)
	if err != nil {
		return nil, err
	}
	_, err = z.Write(content)
	if err != nil {
		return nil, err
	}
	z.Close()

	return &Object{
		Header: header,
		Digest: digest,
		Blob:   b.Bytes(),
	}, nil
}
