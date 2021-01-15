package debug

import (
	"fmt"
	"time"
)

// StopWatch defines the features of a stop-watch must provides.
type StopWatch interface {
	Start(taskName string)
	Stop()
	Duration() time.Duration
	IsRunning() bool
	CurrentTask() string
	PrettyPrint() string
}

// NewStopWatch creates the default StopWatch instance.
func NewStopWatch() StopWatch {
	return &stopWatch{
		tasks:     nil,
		isRunning: false,
	}
}

// stopWatch implements a simple StopWatch but not concurrent-safe, can not be shared by different goroutines.
type stopWatch struct {
	tasks     []string
	durations []time.Duration
	start     time.Time
	end       time.Time
	isRunning bool
}

func (s *stopWatch) Start(taskName string) {
	s.tasks = append(s.tasks, taskName)
	s.start = time.Now()
	s.isRunning = true
}

func (s *stopWatch) Stop() {
	now := time.Now()
	d := now.Sub(s.start)
	s.durations = append(s.durations, d)
	s.isRunning = false
}

func (s *stopWatch) Duration() time.Duration {
	if len(s.durations) == 0 {
		return time.Duration(0)
	}
	return s.durations[len(s.durations)-1]
}

func (s *stopWatch) IsRunning() bool {
	return s.isRunning
}

func (s *stopWatch) CurrentTask() string {
	if len(s.tasks) == 0 {
		return ""
	}
	return s.tasks[len(s.tasks)-1]
}

func (s *stopWatch) PrettyPrint() string {
	header := fmt.Sprintln("------------------")
	cols := fmt.Sprintln("ms\t Task name\t")
	footer := fmt.Sprintln("------------------")

	f := "%v\t %s\t"
	res := header + cols
	for i := range s.tasks {
		res += fmt.Sprintf(f, s.durations[i].Milliseconds(), s.tasks[i]) + "\n"
	}
	res += footer

	return res
}
