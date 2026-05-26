package ansi

import (
	"io"
)

// A LinkElement is used to render hyperlinks.
type LinkElement struct {
	BaseURL  string
	URL      string
	Children []ElementRenderer
	SkipText bool
	SkipHref bool

	hyperlink, resetHyperlink string
	validURL                  bool
}

// Render renders a LinkElement.
func (e *LinkElement) Render(w io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	// Make OSC 8 hyperlink token.
	return nil
}

func (e *LinkElement) renderTextPart(w io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:nestif

func (e *LinkElement) renderHrefPart(w io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}

// makeHyperlink takes a URL and returns an OSC 8 hyperlink token.
func makeHyperlink(link string) (string, string, bool) {
	_ = "STUB: not implemented"
	// Make OSC 8 hyperlink token.
	return "", "", false
}

// if the URL only consists of an anchor, ignore it
