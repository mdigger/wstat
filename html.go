package wstat

import (
	"bytes"
	"errors"
	"io"

	"golang.org/x/net/html"
)

// IgnoreHTMLTags contains the list of names of the HTML tags, the contents of
// which are ignored.
var IgnoreHTMLTags = map[string]struct{}{
	"script": {},
	"style":  {},
	"head":   {},
	"title":  {},
}

// FromHTML extracts text from HTML and returns statistical information on text.
// The contents of the tag from the IgnoreHTMLTAGS list is ignored.
func FromHTML(r io.Reader) (c Counter, err error) {
	var ignoreDepth int
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			err = z.Err()
			if errors.Is(err, io.EOF) {
				err = nil
			}
			return

		case html.TextToken:
			if ignoreDepth > 0 {
				continue
			}

			text := z.Text()
			if len(bytes.TrimSpace(text)) > 0 {
				_, _ = c.Write(text) // only not empty text
			}

		case html.StartTagToken, html.EndTagToken:
			name, _ := z.TagName()
			if _, ok := IgnoreHTMLTags[string(name)]; ok {
				if tt == html.StartTagToken {
					ignoreDepth++
				} else {
					ignoreDepth--
				}
			}
		}
	}
}
