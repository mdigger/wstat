Simple text statistic library
=============================

[![Go Reference](https://pkg.go.dev/badge/github.com/mdigger/wstat.svg)](https://pkg.go.dev/github.com/mdigger/wstat)

Library for quick counting the simplest statistics on the text:
- the total number of characters,
- the number of spaces and separators,
- the number of punctuation symbols,
- the number of digits,
- the number of words,
- allows you to perform calculations in the stream or addition of individual lines,
- supports selection of text from HTML
- counting the number of machine-visiting (typewritten) pages and author's pages
- Calculation of reading time for different languages and different speeds

```golang
stat := wstat.FromString(`text data`)
fmt.Println("reading time:", stat)
```