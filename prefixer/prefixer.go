// Prefixer offers an io.Writer implementation which augments every line with a dynamic prefix.
// The main advantage of this compared to eg the log package is that there is no buffering involved, meaning that every
// call to Write() immediately results in a call to Write() of the underlying writer. This is particularly handy when
// augmenting the output of an external application as it might print important status information without a trailing
// newline.
// This package is not concurrency safe.
package prefixer

import (
	"bytes"
	"fmt"
	"io"
)

type Prefixer struct {
	prefixFunc      func() string
	writer          io.Writer
	trailingNewline bool
	buf             bytes.Buffer // reuse buffer to save allocations
}

// New creates a new Prefixer that forwards all calls to Write() to writer.Write() with all lines prefixed with the
// return value of prefixFunc. Having a function instead of a static prefix allows to print timestamps or other changing
// information.
func New(writer io.Writer, prefixFunc func() string) *Prefixer {
	return &Prefixer{prefixFunc: prefixFunc, writer: writer, trailingNewline: true}
}

func (pf *Prefixer) Write(payload []byte) (int, error) {
	pf.buf.Reset() // clear the buffer

	for _, b := range payload {
		if pf.trailingNewline {
			pf.buf.WriteString(pf.prefixFunc())
			pf.trailingNewline = false
		}

		pf.buf.WriteByte(b)

		if b == '\n' {
			// do not print the prefix right after the newline character as this might
			// be the very last character of the stream and we want to avoid a trailing prefix.
			pf.trailingNewline = true
		}
	}

	n, err := pf.writer.Write(pf.buf.Bytes())
	if err != nil {
		// never return more than original length to satisfy io.Writer interface
		if n > len(payload) {
			n = len(payload)
		}
		return n, err
	}

	// return original length to satisfy io.Writer interface
	return len(payload), nil
}

// EnsureNewline prints a newline if the last character written wasn't a newline unless nothing has ever been written.
// The purpose of this method is to avoid ending the output in the middle of the line.
func (pf *Prefixer) EnsureNewline() {
	if !pf.trailingNewline {
		fmt.Fprintln(pf.writer)
	}
}
