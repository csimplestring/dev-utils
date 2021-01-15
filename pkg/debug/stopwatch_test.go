package debug

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_stopWatch(t *testing.T) {
	s := NewStopWatch()

	assert.False(t, s.IsRunning())
	assert.Equal(t, time.Duration(0), s.Duration())
	assert.Empty(t, s.CurrentTask())

	s.Start("task-1")
	time.Sleep(time.Millisecond * 600)
	assert.True(t, s.IsRunning())
	assert.Equal(t, "task-1", s.CurrentTask())
	s.Stop()

	assert.False(t, s.IsRunning())
	assert.GreaterOrEqual(t, s.Duration(), time.Millisecond*600)
	assert.LessOrEqual(t, s.Duration(), time.Millisecond*610)

	s.Start("task-2")
	time.Sleep(time.Millisecond * 700)
	assert.True(t, s.IsRunning())
	assert.Equal(t, "task-2", s.CurrentTask())
	s.Stop()

	assert.False(t, s.IsRunning())
	assert.GreaterOrEqual(t, s.Duration(), time.Millisecond*700)
	assert.LessOrEqual(t, s.Duration(), time.Millisecond*710)

	t.Log(s.PrettyPrint())
}
