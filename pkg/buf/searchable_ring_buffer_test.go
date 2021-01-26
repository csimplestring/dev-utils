package buf

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_naiveRingBuffer_Match(t *testing.T) {
	r := NewNaiveSearchRingBuffer([]byte{'\r', '0', 38, 48})

	s := []byte{
		'\r', '0', 38, 58, 58, 83, 78, '\r', '0', 38, 48, 83, 92, 78, 4, 38,
	}

	counter := -1
	for _, b := range s {
		counter++
		matched := r.Match(b)
		if counter == 10 {
			assert.True(t, matched, "the pattern should be matched at position 10")
		} else {
			assert.False(t, matched)
		}
	}
}

func TestMatch(t *testing.T) {
	s := []byte{
		'\r', '0', 38, 58, 58, 83, 78, '\r', '0', 38, 48, 83, 92, 78, 4, 38, '\r', '0', 38, 48,
	}
	r := bytes.NewReader(s)

	matched, err := Match(r, []byte{'\r', '0', 38, 48})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(matched))
	assert.Equal(t, []int64{10, 19}, matched)
}
