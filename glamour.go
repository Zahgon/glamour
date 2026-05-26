// Package glamour lets you render markdown documents & templates on ANSI
// compatible terminals. You can create your own stylesheet or simply use one of
// the stylish defaults
package glamour

import (
	"bytes"

	"github.com/yuin/goldmark"

	"charm.land/glamour/v2/ansi"
)

const (
	defaultWidth = 80
	highPriority = 1000
)

// A TermRendererOption sets an option on a TermRenderer.
type TermRendererOption func(*TermRenderer) error

// TermRenderer can be used to render markdown content, posing a depth of
// customization and styles to fit your needs.
type TermRenderer struct {
	md          goldmark.Markdown
	ansiOptions ansi.Options
	buf         bytes.Buffer
	renderBuf   bytes.Buffer
}

// Render initializes a new TermRenderer and renders a markdown with a specific
// style.
func Render(in string, stylePath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// RenderWithEnvironmentConfig initializes a new TermRenderer and renders a
// markdown with a specific style defined by the GLAMOUR_STYLE environment variable.
func RenderWithEnvironmentConfig(in string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RenderBytes initializes a new TermRenderer and renders a markdown with a
// specific style.
func RenderBytes(in []byte, stylePath string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewTermRenderer returns a new TermRenderer the given options.
func NewTermRenderer(options ...TermRendererOption) (*TermRenderer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithBaseURL sets a TermRenderer's base URL.
func WithBaseURL(baseURL string) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithStandardStyle sets a TermRenderer's styles with a standard (builtin)
// style.
func WithStandardStyle(style string) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithEnvironmentConfig sets a TermRenderer's styles based on the
// GLAMOUR_STYLE environment variable.
func WithEnvironmentConfig() TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithStylePath sets a TermRenderer's style from stylePath. stylePath is first
// interpreted as a filename. If no such file exists, it is re-interpreted as a
// standard style.
func WithStylePath(stylePath string) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithStyles sets a TermRenderer's styles.
func WithStyles(styles ansi.StyleConfig) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithStylesFromJSONBytes sets a TermRenderer's styles by parsing styles from
// jsonBytes.
func WithStylesFromJSONBytes(jsonBytes []byte) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithStylesFromJSONFile sets a TermRenderer's styles from a JSON file.
func WithStylesFromJSONFile(filename string) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithWordWrap sets a TermRenderer's word wrap.
func WithWordWrap(wordWrap int) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithTableWrap controls whether table content will wrap if too long.
// This is true by default. If false, table content will be truncated with an
// ellipsis if too long to fit.
func WithTableWrap(tableWrap bool) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithInlineTableLinks forces tables to render links inline. By default,links
// are rendered as a list of links at the bottom of the table.
func WithInlineTableLinks(inlineTableLinks bool) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithPreservedNewLines preserves newlines from being replaced.
func WithPreservedNewLines() TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithEmoji sets a TermRenderer's emoji rendering.
func WithEmoji() TermRendererOption { _ = "STUB: not implemented"; return *new(TermRendererOption) }

// WithChromaFormatter sets a TermRenderer's chroma formatter used for code blocks.
func WithChromaFormatter(formatter string) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

// WithOptions sets multiple TermRenderer options within a single TermRendererOption.
func WithOptions(options ...TermRendererOption) TermRendererOption {
	_ = "STUB: not implemented"
	return *new(TermRendererOption)
}

func (tr *TermRenderer) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (tr *TermRenderer) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close must be called after writing to TermRenderer. You can then retrieve
// the rendered markdown by calling Read.
func (tr *TermRenderer) Close() error { _ = "STUB: not implemented"; return nil }

// Render returns the markdown rendered into a string.
func (tr *TermRenderer) Render(in string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RenderBytes returns the markdown rendered into a byte slice.
func (tr *TermRenderer) RenderBytes(in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getEnvironmentStyle() string { _ = "STUB: not implemented"; return "" }

func getDefaultStyle(style string) (*ansi.StyleConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
