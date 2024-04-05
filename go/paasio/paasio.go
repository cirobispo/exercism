package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	rmx         sync.Mutex
	reader      io.Reader
	rBytesCount int64
	rCallsCount int
}

type writeCounter struct {
	wmx         sync.Mutex
	writer      io.Writer
	wBytesCount int64
	wCallsCount int
}

// For the return of the function NewReadWriteCounter,
// you must also define a type that satisfies the ReadWriteCounter interface.

type readWriteCounter struct {
	readCounter
	writeCounter
}

func (rc *readCounter) Read(p []byte) (int, error) {
	var count int
	if c, err := rc.reader.Read(p); err != nil {
		return count, err
	} else {
		count = c
	}

	defer rc.rmx.Unlock()
	rc.rmx.Lock()
	rc.rBytesCount += int64(count)
	rc.rCallsCount++

	return count, nil
}

func (rc *readCounter) ReadCount() (int64, int) {
	defer rc.rmx.Unlock()
	rc.rmx.Lock()
	bc, cc := rc.rBytesCount, rc.rCallsCount

	return bc, cc
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	var count int
	if c, err := wc.writer.Write(p); err != nil {
		return c, err
	} else {
		count = c
	}

	defer wc.wmx.Unlock()
	wc.wmx.Lock()
	wc.wBytesCount += int64(count)
	wc.wCallsCount++

	return count, nil
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	defer wc.wmx.Unlock()
	wc.wmx.Lock()
	bc, cc := wc.wBytesCount, wc.wCallsCount

	return bc, cc
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter:  readCounter{reader: readwriter},
		writeCounter: writeCounter{writer: readwriter},
	}
}
