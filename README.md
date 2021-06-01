Simple text statistic library
=============================

[![Go Reference](https://pkg.go.dev/badge/github.com/mdigger/wstat.svg)](https://pkg.go.dev/github.com/mdigger/wstat)

Library for quick counting the simplest statistics on the text:
- the total number of characters,
- the number of spaces and separators,
- the number of words,
- allows you to perform calculations in the stream or addition of individual lines,
- supports selection of text from HTML
- counting the number of machine-visiting pages and author's pages

```golang
stat := wstat.FromString(sample)
fmt.Printf(`
--- stats -----------
chars:        %v
spaces:       %v
words:        %v
--- pages -----------
typewritten:  %v
author's:     %v
--- reading time ----
duration:     %v
---------------------
`,
    stat.Chars, stat.Spaces, stat.Words,
    stat.Pages(), stat.AuthorPages(), stat.Duration(275))
```