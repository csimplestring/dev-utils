package file

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastModifiedMonitor_GetCurrentState(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpfile.Name())

	tmpFileStat, err := tmpfile.Stat()
	if err != nil {
		t.Error(err)
	}

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		l       *LastModifiedMonitor
		args    args
		want    State
		wantErr bool
	}{
		{
			name:    "no-exist file",
			l:       &LastModifiedMonitor{},
			args:    args{path: "./non-exist.txt"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "temp file",
			l:       &LastModifiedMonitor{},
			args:    args{path: tmpfile.Name()},
			want:    lastModTimeState{lastMod: tmpFileStat.ModTime()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LastModifiedMonitor{}
			got, err := l.GetCurrentState(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastModifiedMonitor.GetCurrentState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastModifiedMonitor.GetCurrentState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_watcher_Check(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpfile.Name())

	w, err := NewWatcher(tmpfile.Name(), &LastModifiedMonitor{})
	assert.NoError(t, err)

	cnt := 0
	cb := func(path string) {
		cnt++
	}

	err = w.Check(cb)
	assert.NoError(t, err)
	assert.Equal(t, 0, cnt)

	_, err = tmpfile.WriteString("add")
	assert.NoError(t, err)

	err = w.Check(cb)
	assert.NoError(t, err)
	assert.Equal(t, 1, cnt)

	err = w.Check(cb)
	assert.NoError(t, err)
	assert.Equal(t, 1, cnt)
}
