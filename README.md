# dev-utils
Try to give a comprehensive list of utilities during Go development. Keep small and elegant. 

[![GitHub license](https://img.shields.io/github/license/csimplestring/dev-utils)](https://github.com/csimplestring/dev-utils/blob/main/LICENSE) [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org) <a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-80%25-brightgreen.svg?longCache=true&style=flat)</a> [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)


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

