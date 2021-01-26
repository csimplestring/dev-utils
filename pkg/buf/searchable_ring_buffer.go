package buf

import "io"

// Matcher ...
type Matcher interface {
	Match(b byte) bool
}

// A RingBuffer that can be used to scan byte sequences for subsequences.
// This class implements an efficient naive search algorithm,
// which allows the user of the library to identify byte sequences in a stream on-the-fly
// so that the stream can be segmented without having to buffer the data.
type naiveRingBuffer struct {
	search       []byte
	buf          []byte
	insertionPos int
	bufSize      int
}

// NewNaiveSearchRingBuffer creates an instance of naiveRingBuffer
func NewNaiveSearchRingBuffer(search []byte) Matcher {
	buf := make([]byte, len(search))
	copy(buf, search)

	return &naiveRingBuffer{
		search:       search,
		buf:          buf,
		insertionPos: 0,
		bufSize:      0,
	}
}

// Add the given byte to the buffer and notify whether or not the byte completes the desired byte sequence.
// return true if this byte completes the byte sequence, false otherwise.
func (n *naiveRingBuffer) Match(b byte) bool {
	n.buf[n.insertionPos] = b
	n.insertionPos = (n.insertionPos + 1) % len(n.search)

	n.bufSize = n.bufSize + 1
	if n.bufSize < len(n.search) {
		return false
	}

	for i, b := range n.search {
		pos := (n.insertionPos + i) % len(n.search)
		if b != n.buf[pos] {
			return false
		}
	}

	return true
}

// Match is an util function that it tries to find the matches in r.
// returns all the matched start positions.
func Match(r io.ByteReader, search []byte) ([]int64, error) {

	ring := NewNaiveSearchRingBuffer(search)

	var pos int64 = 0
	var matchedPos []int64
	for {
		b, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if ring.Match(b) {
			matchedPos = append(matchedPos, pos)
		}
		pos++
	}

	return matchedPos, nil
}
