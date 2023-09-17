package gogit_test

import (
	"encoding/hex"
	"testing"

	"github.com/fardream/gogit"
)

func TestTree_Content(t *testing.T) {
	bak, _ := hex.DecodeString("d8329fc1cc938780ffdd9f94e0d364e0ea74f579")
	new_txt, _ := hex.DecodeString("fa49b077972391ad58037050f2a75f74e3671e92")
	test_txt, _ := hex.DecodeString("1f7a7a472abf3dd9643fd615f6da379c4acb3e3a")

	tree := new(gogit.Tree)
	if err := tree.Add(gogit.Mode_RegularFile, "new.txt", new_txt); err != nil {
		t.Fatal(err)
	}
	if err := tree.Add(gogit.Mode_RegularFile, "test.txt", test_txt); err != nil {
		t.Fatal(err)
	}
	if err := tree.Add(gogit.Mode_Tree, "bak", bak); err != nil {
		t.Fatal(err)
	}

	content, err := tree.Content()
	if err != nil {
		t.Fatal(err)
	}

	obj, err := gogit.NewObject(gogit.ContentType_Tree, content)
	if err != nil {
		t.Fatal(err)
	}

	want := "3c4e9cd789d88d8d89c1073707c3585e41b0e614"
	if obj.HexDigest() != want {
		t.Fatalf("want: %s, got: %s", want, obj.HexDigest())
	}
}
