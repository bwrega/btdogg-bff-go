package loader

import "testing"
import "github.com/stretchr/testify/assert"

func TestShouldLoadMockTorrents(t *testing.T) {
	torrents := Load()
	assert.Len(t, torrents, 10)
}
