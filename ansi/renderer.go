package ansi

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Options is used to configure an ANSIRenderer.
type Options struct {
	BaseURL          string
	WordWrap         int
	TableWrap        *bool
	InlineTableLinks bool
	PreserveNewLines bool
	Styles           StyleConfig
	ChromaFormatter  string
}

// ANSIRenderer renders markdown content as ANSI escaped sequences.
type ANSIRenderer struct { //nolint: revive
	context RenderContext
}

// NewRenderer returns a new ANSIRenderer with style and options set.
func NewRenderer(options Options) *ANSIRenderer { _ = "STUB: not implemented"; return nil }

// RegisterFuncs implements NodeRenderer.RegisterFuncs.
func (r *ANSIRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	_ = "STUB: not implemented"
	// blocks
	return
}

// inlines

// tables

// definitions

// footnotes

// checkboxes

// strikethrough

// emoji

func (r *ANSIRenderer) renderNode(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

// children get rendered by their parent

//nolint: nestif
// everything below the Document element gets rendered into a block buffer

// everything below the Document element gets rendered into a block buffer

// if we're finished rendering the entire document,
// flush to the real writer

func isChild(node ast.Node) bool { _ = "STUB: not implemented"; return false }

// These types are already rendered by their parent

func resolveRelativeURL(baseURL string, rel string) string { _ = "STUB: not implemented"; return "" }
