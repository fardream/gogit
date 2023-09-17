package gogit

import (
	"bytes"
	"fmt"
	"io"
	"slices"
)

type TreeEntry struct {
	Mode int32  `json:"mode,omitempty"`
	Name string `json:"name,omitempty"`
	Sha1 []byte `json:"sha1,omitempty"`
}

type Tree struct {
	Entries []*TreeEntry
}

// NewTreeEntry creates a new [TreeEntry].
func NewTreeEntry(mode Mode, name string, sha1 []byte) *TreeEntry {
	return &TreeEntry{
		Mode: int32(mode),
		Name: name,
		Sha1: sha1,
	}
}

// WriteContent writes the content of the tree to the given [io.Writer].
func (tree *Tree) WriteContent(w io.Writer) error {
	for _, e := range tree.Entries {
		mode := Mode(e.Mode)
		_, err := fmt.Fprintf(w, "%s %s", mode.String(), e.Name)
		if err != nil {
			return fmt.Errorf("failed to write entry %s: %w", e.Name, err)
		}
		_, err = w.Write([]byte{0})
		if err != nil {
			return fmt.Errorf("failed to write 0 seperate for name/sha1 for %s: %w", e.Name, err)
		}
		_, err = w.Write(e.Sha1)
		if err != nil {
			return fmt.Errorf("failed to write sha1 for %s: %w", e.Name, err)
		}

	}

	return nil
}

// Content generates the byte array (uncompressed for tree)
func (tree *Tree) Content() ([]byte, error) {
	var b bytes.Buffer

	if err := tree.WriteContent(&b); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// findName finds index of the entry by name using the [slices.IndexFunc] function. returns -1 if not found.
func (tree *Tree) findName(name string) int {
	return slices.IndexFunc(tree.Entries, func(e *TreeEntry) bool {
		return e.Name == name
	})
}

// compareTreeEntryByName can be supplied to [slices.SortFunc] and sort the tree entry list.
func compareTreeEntryByName(l, r *TreeEntry) int {
	switch {
	case l.Name < r.Name:
		return -1
	case l.Name > r.Name:
		return 1
	default:
		return 0
	}
}

// Add adds an entry to the tree, error if it's a duplicate
func (tree *Tree) Add(mode Mode, name string, sha1 []byte) error {
	if tree.findName(name) >= 0 {
		return fmt.Errorf("name %s already exists in the tree, use replace instead", name)
	}
	entry := &TreeEntry{
		Mode: int32(mode),
		Name: name,
		Sha1: sha1,
	}

	tree.Entries = append(tree.Entries, entry)

	slices.SortFunc(
		tree.Entries, compareTreeEntryByName,
	)

	return nil
}

// Replace an entry in the tree by its name, error if not found.
func (tree *Tree) Replace(mode Mode, name string, sha1 []byte) error {
	idx := tree.findName(name)
	if idx < 0 {
		return fmt.Errorf("name %s doesn't exist in the tree, use add instead", name)
	}

	tree.Entries[idx] = &TreeEntry{
		Mode: int32(mode),
		Name: name,
		Sha1: sha1,
	}

	return nil
}

// AddOrReplace will replace an entry if the name already exists, or add it to tree if the name doesn't exist.
func (tree *Tree) AddOrReplace(mode Mode, name string, sha1 []byte) error {
	entry := &TreeEntry{
		Mode: int32(mode),
		Name: name,
		Sha1: sha1,
	}

	idx := tree.findName(name)

	if idx < 0 {
		tree.Entries = append(tree.Entries, entry)
		slices.SortFunc(tree.Entries, compareTreeEntryByName)
	} else {
		tree.Entries[idx] = entry
	}

	return nil
}
