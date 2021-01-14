# dev-utils
Try to give a comprehensive list of utilities during Go development. Keep small and elegant. 

## install
```
go get github.com/csimplestring/dev-utils
```

# features
- file watcher

# examples
- *file watcher*: The file watch will detects if the file is changed or not, and execute a callback if it is changed. 
```
    // In this example, the last-modified-time determines if a file is modified or not.
    // if you want to use the md5 sum of a file content as the *state*, you can create your own Monitor to detect if the file state is changed by implementing the UpdateModifer interface.

    w, err := NewWatcher(tmpfile.Name(), &LastModifiedMonitor{})
	assert.NoError(t, err)

	cnt := 0
	cb := func(path string) {
		cnt++
	}

    // file is not changed, cnt == 0
	err = w.Check(cb)
	assert.NoError(t, err)
	assert.Equal(t, 0, cnt)

	_, err = tmpfile.WriteString("add")
	assert.NoError(t, err)

    // file is changed, last modified time is updated, cnt == 1
	err = w.Check(cb)
	assert.NoError(t, err)
	assert.Equal(t, 1, cnt)

``` 

