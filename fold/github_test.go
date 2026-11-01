package fold

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGithubFolder(t *testing.T) {
	var b bytes.Buffer
	f := NewGithub(&b)
	id, err := f.Start("lint a/b", "hello", true)
	assert.NoError(t, err)
	assert.Empty(t, id)
	err = f.Stop(id)
	assert.NoError(t, err)
	assert.Equal(t, "::group::hello lint a/b\n::endgroup::\n", b.String(), "it should only group, not annotate")
}
