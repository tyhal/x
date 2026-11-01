package fold

import (
	"fmt"
	"io"
)

type githubFolder struct {
	io.Writer
}

// NewGithub uses the GitHub Actions log syntax
func NewGithub(w io.Writer) Folder {
	return &githubFolder{w}
}

// Start only groups, annotating a file (::error file=...) is left to whatever knows its path and position
func (g githubFolder) Start(title string, msg string, _ bool) (string, error) {
	_, err := fmt.Fprintf(g, "::group::%s %s\n", msg, title)
	return "", err
}

func (g githubFolder) Stop(_ string) error {
	_, err := fmt.Fprintf(g, "::endgroup::\n")
	return err
}
