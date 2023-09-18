package gogit

import (
	"encoding/hex"
	"fmt"
)

const sha1Length = 20

type Sha1 [sha1Length]byte

func (s *Sha1) String() string {
	return hex.EncodeToString(s[:])
}

func (s Sha1) MarshalText() (text []byte, err error) {
	text = make([]byte, sha1Length*2)
	hex.Encode(text, s[:])
	return
}

func (s *Sha1) UnmarshalText(text []byte) error {
	if len(text) < sha1Length*2 {
		return fmt.Errorf("text is too small %d for sha1 digest", len(text))
	}

	n, err := hex.Decode((*s)[:], text)
	if err != nil {
		return err
	}
	if n != sha1Length {
		return fmt.Errorf("data is incomplete, read %d", n)
	}

	return nil
}

func (s *Sha1) Equals(e *Sha1) bool {
	for i := 0; i < sha1Length; i++ {
		if s[i] != e[i] {
			return false
		}
	}

	return true
}

func (s *Sha1) UnmarshalString(src string) error {
	return s.UnmarshalText([]byte(src))
}
