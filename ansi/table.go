package ansi

import (
	"io"

	"charm.land/lipgloss/v2/table"
	astext "github.com/yuin/goldmark/extension/ast"
)

// A TableElement is used to render tables.
type TableElement struct {
	lipgloss *table.Table
	table    *astext.Table
	header   []string
	row      []string
	source   []byte

	tableImages []tableLink
	tableLinks  []tableLink
}

// A TableRowElement is used to render a single row in a table.
type TableRowElement struct{}

// A TableHeadElement is used to render a table's head element.
type TableHeadElement struct{}

// A TableCellElement is used to render a single cell in a row.
type TableCellElement struct {
	Children []ElementRenderer
	Head     bool
}

// Render renders a TableElement.
func (e *TableElement) Render(w io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

//nolint:errcheck

//nolint: gosec

func (e *TableElement) setStyles(ctx RenderContext) { _ = "STUB: not implemented"; return }

// Default Styles

// Override with custom styles

//nolint: gosec

// do nothing

func (e *TableElement) setBorders(ctx RenderContext) { _ = "STUB: not implemented"; return }

// Finish finishes rendering a TableElement.
func (e *TableElement) Finish(_ io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Finish finishes rendering a TableRowElement.
func (e *TableRowElement) Finish(_ io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Finish finishes rendering a TableHeadElement.
func (e *TableHeadElement) Finish(_ io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Render renders a TableCellElement.
func (e *TableCellElement) Render(_ io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}
