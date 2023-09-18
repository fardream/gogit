package gogit_test

import (
	"testing"

	"github.com/fardream/gogit"
)

func TestCommit_Content(t *testing.T) {
	commit := &gogit.Commit{}

	var parent gogit.Sha1

	commit.Tree.UnmarshalString("3c4e9cd789d88d8d89c1073707c3585e41b0e614")
	parent.UnmarshalString("ace9c3adeade6b88c9bb4755b8670cdaaa5757a2")
	commit.Parents = append(commit.Parents, parent)
	commit.SetAuthor("Chao Xu", "fardream@users.noreply.github.com", 1694803547, "-0400")
	commit.SetCommittor("Chao Xu", "fardream@users.noreply.github.com", 1694803547, "-0400")
	commit.CommitMessage = "Third commit\n"

	content, _ := commit.Content()
	o, err := gogit.NewObject(gogit.ContentType_Commit, content)
	if err != nil {
		t.Fatal(err)
	}
	want := "10ed2c5136df4c7a0f78d923d10a8fab342335fb"
	if o.HexDigest() != want {
		t.Fatalf("want: %s, got: %s", want, o.HexDigest())
	}
}
