package htmlparser

import (
	"io"
)

// ReadableText reperisents a HTML node (p, hX, code, pre)
type ReadableText struct {
	ContentType string
	Content     string
}

type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML doc and return a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
