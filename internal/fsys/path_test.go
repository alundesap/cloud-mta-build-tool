package dir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPath(t *testing.T) {
	t.Parallel()
	path, err := GetCurrentPath()
	assert.NotEmpty(t, path)
	assert.Nil(t, err)
}