package gogit_test

import (
	"testing"

	"github.com/fardream/gogit"
)

func TestNewBlob(t *testing.T) {
	content := "test content\n"
	blob, err := gogit.NewObject(gogit.ContentType_Blob, []byte(content))
	if err != nil {
		t.Fatalf("failed to generate blob: %s", err.Error())
	}
	expectedHash := "d670460b4b4aece5915caf5c68d12f560a9fe3e4"
	if blob.HexDigest() != expectedHash {
		t.Errorf("expecting hash %s, got %s", expectedHash, blob.HexDigest())
	}
}
