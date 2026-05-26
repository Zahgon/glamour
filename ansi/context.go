package ansi

import (
	"github.com/microcosm-cc/bluemonday"
)

// RenderContext holds the current rendering options and state.
type RenderContext struct {
	options Options

	blockStack *BlockStack
	table      *TableElement

	stripper *bluemonday.Policy
}

// NewRenderContext returns a new RenderContext.
func NewRenderContext(options Options) RenderContext {
	_ = "STUB: not implemented"
	return *new(RenderContext)
}

// SanitizeHTML sanitizes HTML content.
func (ctx RenderContext) SanitizeHTML(s string, trimSpaces bool) string {
	_ = "STUB: not implemented"
	return ""
}
