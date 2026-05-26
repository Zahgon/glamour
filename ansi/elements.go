package ansi

import (
	"io"

	"github.com/yuin/goldmark/ast"
)

// ElementRenderer is called when entering a markdown node.
type ElementRenderer interface {
	Render(w io.Writer, ctx RenderContext) error
}

// StyleOverriderElementRenderer is called when entering a markdown node with a specific style.
type StyleOverriderElementRenderer interface {
	StyleOverrideRender(w io.Writer, ctx RenderContext, style StylePrimitive) error
}

// ElementFinisher is called when leaving a markdown node.
type ElementFinisher interface {
	Finish(w io.Writer, ctx RenderContext) error
}

// An Element is used to instruct the renderer how to handle individual markdown
// nodes.
type Element struct {
	Entering string
	Exiting  string
	Renderer ElementRenderer
	Finisher ElementFinisher
}

// NewElement returns the appropriate render Element for a given node.
func (tr *ANSIRenderer) NewElement(node ast.Node, source []byte) Element {
	_ = "STUB: not implemented"
	return *new(Element)
}

// Document

// Heading

// Paragraph

// Blockquote

// Lists

//nolint: gosec

// Text Elements

//nolint: staticcheck

// Links

// For non-email links, skip text (show only href)
// For email links, skip href (hide mailto: URL)

// Images

//nolint: staticcheck

// Code

//nolint: staticcheck

// Tables

// HTML Elements

//nolint: staticcheck

//nolint: staticcheck

// Definition Lists

// Handled by parents

// handled by KindListItem

// Unknown case
