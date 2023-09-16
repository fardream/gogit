package gogit

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
)

func NewBlob(content []byte) (*Blob, error) {
	l := len(content)
	header, err := Header(ContentType_Blob, l)
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
	digestBytes := hasher.Sum(nil)
	digest := hex.EncodeToString(digestBytes)
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

	return &Blob{
		Header:    header,
		HexDigest: digest,
		Blob:      b.Bytes(),
	}, nil
}
