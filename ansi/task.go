package ansi

import (
	"io"
)

// A TaskElement is used to render tasks inside a todo-list.
type TaskElement struct {
	Checked bool
}

// Render renders a TaskElement.
func (e *TaskElement) Render(w io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}
