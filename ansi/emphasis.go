package ansi

import (
	"io"
)

// A EmphasisElement is used to render emphasis.
type EmphasisElement struct {
	Children []ElementRenderer
	Level    int
}

// Render renders a EmphasisElement.
func (e *EmphasisElement) Render(w io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}

// StyleOverrideRender renders a EmphasisElement with a given style.
func (e *EmphasisElement) StyleOverrideRender(w io.Writer, ctx RenderContext, style StylePrimitive) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *EmphasisElement) doRender(w io.Writer, ctx RenderContext, style StylePrimitive) error {
	_ = "STUB: not implemented"
	return nil
}
