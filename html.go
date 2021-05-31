package wstat

import (
	"errors"
	"io"

	"golang.org/x/net/html"
)

var IgnoreHTMLTags = map[string]struct{}{
	"script": {},
	"style":  {},
	"head":   {},
	"title":  {},
}

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
			if ignoreDepth == 0 {
				_, _ = c.Write(z.Text())
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
