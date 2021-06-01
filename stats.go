// Library for quick counting the simplest statistics on the text:
package wstat

import (
	"bufio"
	"errors"
	"io"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	// One typewritten page accommodates 1860 printed signs (including spaces)
	PageChars = 1_860
	// In the USSR and the Russian Federation, an account of the author's leaf
	// is equal to 40,000 printed signs (including punctuation marks, numbers
	// and spaces between words to the fields)
	AuthorChars = 40_000
)

// Counter contains data with statistical counting results.
type Counter struct {
	Chars  int `json:"chars"`  // Total number of characters
	Spaces int `json:"spaces"` // Number of spaces and separators
	Words  int `json:"words"`  // Number of words
}

// ReadFrom calc and add statistical information about the text from stream.
// io.ReaderFrom interface support.
func (c *Counter) ReadFrom(r io.Reader) (n int64, err error) {
	var reader io.RuneReader
	if rr, ok := r.(io.RuneReader); ok {
		reader = rr
	} else {
		reader = bufio.NewReader(r)
	}

	var isInWord bool
	for {
		r, l, err := reader.ReadRune()
		n += int64(l)

		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}
			return n, err
		}

		isInWord = c.writeRune(r, isInWord)
	}
}

// Write allows you to transfer a set of bytes to account for text statistics.
// Supports an io.Writer interface.
func (c *Counter) Write(s []byte) (n int, err error) {
	n = len(s)
	var isInWord bool
	for len(s) > 0 {
		r, l := utf8.DecodeRune(s)
		isInWord = c.writeRune(r, isInWord)
		s = s[l:]
	}

	return n, nil
}

// Write allows you to transfer a strings to account for text statistics.
// Supports an io.StringWriter interface.
func (c *Counter) WriteString(s string) (n int, err error) {
	var isInWord bool
	for _, r := range s {
		isInWord = c.writeRune(r, isInWord)
	}

	return len(s), nil
}

func (c *Counter) writeRune(r rune, isInWord bool) bool {
	c.Chars++

	switch {
	case unicode.IsSpace(r):
		c.Spaces++
		return false
	case unicode.IsPunct(r):
		return false
	case unicode.IsLetter(r):
		if !isInWord {
			c.Words++
		}
		return true
	default:
		return isInWord
	}
}

// Pages returns an approximate number of standard typewritten pages,
// which takes the processed text.
func (c Counter) Pages() int {
	return (c.Chars + PageChars - 1) / PageChars
}

// AuthorPages returns the number of author pages that occupy processed text.
func (c Counter) AuthorPages() float32 {
	return float32(c.Chars) / float32(AuthorChars)
}

// Duration will return the approximate text reading time at a given speed
// reading (words per minute).
//
// The average speed (words per minute):
// 	English — 275 (reading time), 180 (speaking time)
// 	Russian — 200 (reading time), 130 (speaking time).
func (c Counter) Duration(wps int) time.Duration {
	return (time.Duration(c.Words) * time.Minute / time.Duration(wps)).
		Round(time.Second)
}

// Join returns a new statistics counter with total data.
func (c Counter) Join(counters ...Counter) Counter {
	for _, c2 := range counters {
		c.Chars += c2.Chars
		c.Spaces += c2.Spaces
		c.Words += c2.Words
	}

	return c
}

// ReadFrom returns statistical information about the text from stream.
func ReadFrom(r io.Reader) (c Counter, err error) {
	_, err = c.ReadFrom(r)
	return
}

// FromString returns statistical information about the text from string.
func FromString(s string) (c Counter) {
	_, _ = c.WriteString(s)
	return
}

// FromBytes returns statistical information about the text from bytes.
func FromBytes(b []byte) (c Counter) {
	_, _ = c.Write(b)
	return
}
