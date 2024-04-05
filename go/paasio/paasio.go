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
	count := 0
	var err error
	if count, err = rc.reader.Read(p); err != nil {
		return count, err
	}

	rc.rmx.Lock()
	rc.rBytesCount += int64(count)
	rc.rCallsCount++
	rc.rmx.Unlock()

	return count, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.rmx.Lock()
	bc, cc := rc.rBytesCount, rc.rCallsCount
	rc.rmx.Unlock()

	return bc, cc
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	count := 0
	var err error
	if count, err = wc.writer.Write(p); err != nil {
		return count, err
	}

	wc.wmx.Lock()
	wc.wBytesCount += int64(count)
	wc.wCallsCount++
	wc.wmx.Unlock()

	return count, err
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	wc.wmx.Lock()
	bc, cc := wc.wBytesCount, wc.wCallsCount
	wc.wmx.Unlock()

	return bc, cc
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{rmx: sync.Mutex{}, reader: reader, rBytesCount: 0}
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{wmx: sync.Mutex{}, writer: writer, wBytesCount: 0}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	r := readCounter{reader: readwriter}
	w := writeCounter{writer: readwriter}
	return &readWriteCounter{readCounter: r, writeCounter: w}
}
