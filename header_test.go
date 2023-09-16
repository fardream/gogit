package gogit_test

import (
	"encoding/hex"
	"testing"

	"github.com/fardream/gogit"
)

func TestNewHeader(t *testing.T) {
	h, err := gogit.NewHeader(gogit.ContentType_Blob, 50)
	if err != nil {
		t.Fatalf("failed to generate header: %s", err.Error())
	}

	expected := hex.EncodeToString(append([]byte("blob 50"), 0))
	got := hex.EncodeToString(h)

	if expected != got {
		t.Errorf("expecting: %s, got: %s", expected, got)
	}
}
