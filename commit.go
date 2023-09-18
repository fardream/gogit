package gogit

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

type commitUser struct {
	Name     string
	Email    string
	TimeSecs int
	Offset   string
}

func (u *commitUser) setData(name string, email string, timesecs int, offset string) {
	u.Name = name
	u.Email = email
	u.TimeSecs = timesecs
	u.Offset = offset
}

func (u *commitUser) String() string {
	return fmt.Sprintf("%s <%s> %d %s", u.Name, u.Email, u.TimeSecs, u.Offset)
}

type Commit struct {
	Tree          Sha1
	Parents       []Sha1
	Author        commitUser
	Committor     commitUser
	CommitMessage string
}

func (c *Commit) SetAuthor(name string, email string, timesecs int, offset string) {
	c.Author.setData(name, email, timesecs, offset)
}

func (c *Commit) SetCommittor(name string, email string, timesecs int, offset string) {
	c.Committor.setData(name, email, timesecs, offset)
}

func TimeSecsOffset(t time.Time) (secs int, offset string) {
	offset = t.Format("-0700")
	_, v := t.Zone()
	secs = int(t.Local().Unix()) + v
	return
}

func (c *Commit) WriteContent(w io.Writer) (err error) {
	_, err = fmt.Fprintf(w, "tree %s\n", c.Tree.String())
	if err != nil {
		return
	}
	for _, p := range c.Parents {
		_, err = fmt.Fprintf(w, "parent %s\n", p.String())
		if err != nil {
			return
		}
	}
	_, err = fmt.Fprintf(w, "author %s\n", c.Author.String())
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, "committer %s\n", c.Committor.String())
	if err != nil {
		return
	}

	_, err = fmt.Fprintln(w)
	if err != nil {
		return
	}
	_, err = fmt.Fprint(w, c.CommitMessage)
	if err != nil {
		return
	}

	return
}

func (c *Commit) Content() ([]byte, error) {
	var b bytes.Buffer

	if err := c.WriteContent(&b); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
