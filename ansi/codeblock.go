package ansi

import (
	"io"
	"sync"
)

const (
	// The chroma style theme name used for rendering.
	chromaStyleTheme = "charm"

	// The chroma formatter name used for rendering.
	chromaFormatter = "terminal256"
)

// mutex for synchronizing access to the chroma style registry.
// Related https://github.com/alecthomas/chroma/pull/650
var mutex = sync.Mutex{}

// A CodeBlockElement is used to render code blocks.
type CodeBlockElement struct {
	Code     string
	Language string
}

func chromaStyle(style StylePrimitive) string { _ = "STUB: not implemented"; return "" }

// Render renders a CodeBlockElement.
func (e *CodeBlockElement) Render(w io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Don't register the style if it's already registered.

//nolint:gosec

//nolint:errcheck

// fallback rendering
