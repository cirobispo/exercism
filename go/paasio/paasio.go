package paasio

import "io"

// Define readCounter and writeCounter types here.
type readCounter struct {
	reader      io.Reader
	rBytesCount int64
	rCallsCount int
}

type writeCounter struct {
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
	rc.rBytesCount += int64(count)
	rc.rCallsCount++

	return count, err
}

func (rc readCounter) ReadCount() (int64, int) {
	return rc.rBytesCount, rc.rCallsCount
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	count := 0
	var err error
	if count, err = wc.writer.Write(p); err != nil {
		return count, err
	}
	wc.wBytesCount += int64(count)
	wc.wCallsCount++

	return count, err
}

func (wc writeCounter) WriteCount() (n int64, nops int) {
	return wc.wBytesCount, wc.wCallsCount
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader, rBytesCount: 0}
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer, wBytesCount: 0}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	r := readCounter{reader: readwriter}
	w := writeCounter{writer: readwriter}
	return &readWriteCounter{readCounter: r, writeCounter: w}
}
