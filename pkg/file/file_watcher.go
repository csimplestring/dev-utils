package file

import (
	"os"
	"sync"
	"time"
)

// State represents the state of a file to watch, can be lastModifiedTime, md5 sum eetc.
type State interface {
	Equal(s State) bool
}

// lastModTimeState implements State interface, uses the lastModifiedTime of a file as the state to watch
type lastModTimeState struct {
	lastMod time.Time
}

// Equal compares the lastMod field
func (l lastModTimeState) Equal(s State) bool {
	if other, ok := s.(lastModTimeState); ok {
		return other.lastMod.Equal(l.lastMod)
	}

	return false
}

// UpdateMonitor defines the monitor, given a path of file, returns the current State.
type UpdateMonitor interface {
	GetCurrentState(path string) (State, error)
}

// LastModifiedMonitor implements UpdateMonitor interface, gets the last modified time of file and returns it as State.
type LastModifiedMonitor struct {
}

// GetCurrentState returns the current state which contains file last modified time.
func (l *LastModifiedMonitor) GetCurrentState(path string) (State, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return lastModTimeState{lastMod: info.ModTime()}, nil
}

// Watcher allows the user to configure a path to watch for modification and potentially poll to check if thee file has
// been changed. If the file's state is modified, the callback function will be executed.
type Watcher interface {
	Check(callback func(path string)) error
}

// watcher gives a default synchrounous implementation of Watch, it is safe to be called by goroutines.
// the user can extend and pass the UpdateMonitor to achieve different purpose.
type watcher struct {
	monitor   UpdateMonitor
	m         sync.Mutex
	lastState State
	path      string
}

// NewWatcher creates a new watcher instance.
func NewWatcher(path string, monitor UpdateMonitor) (Watcher, error) {
	c, err := monitor.GetCurrentState(path)
	if err != nil {
		return nil, err
	}

	return &watcher{
		monitor:   monitor,
		m:         sync.Mutex{},
		lastState: c,
		path:      path,
	}, err
}

// Check will trigger the callback function when it detects the state is changed.
func (w *watcher) Check(callback func(path string)) error {
	w.m.Lock()
	defer w.m.Unlock()

	noState := w.lastState == nil

	c, err := w.monitor.GetCurrentState(w.path)
	if err != nil {
		return err
	}

	if noState {
		w.lastState = c
		return nil
	}

	if c.Equal(w.lastState) {
		return nil
	}

	w.lastState = c
	callback(w.path)
	return nil
}
