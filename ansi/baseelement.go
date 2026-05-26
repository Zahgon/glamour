package ansi

import (
	"io"
	"strings"
)

// BaseElement renders a styled primitive element.
type BaseElement struct {
	Token  string
	Prefix string
	Suffix string
	Style  StylePrimitive
}

func formatToken(format string, token string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func renderText(w io.Writer, rules StylePrimitive, s string) (int, error) {
	_ = "STUB: not implemented" //nolint:unparam
	return 0, nil
}

// XXX: We're using [ansi.Style] instead of [lipgloss.Style] because
// Lip Gloss has a weird bug where it adds spaces when rendering joined
// strings. Needs further investigation.

// StyleOverrideRender renders a BaseElement with an overridden style.
func (e *BaseElement) StyleOverrideRender(w io.Writer, ctx RenderContext, style StylePrimitive) error {
	_ = "STUB: not implemented"
	return nil
}

// Render renders a BaseElement.
func (e *BaseElement) Render(w io.Writer, ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *BaseElement) doRender(w io.Writer, st1, st2 StylePrimitive) error {
	_ = "STUB: not implemented"
	return nil
}

// render unstyled prefix/suffix

// render styled prefix/suffix

// https://www.markdownguide.org/basic-syntax/#characters-you-can-escape
var escapeReplacer = strings.NewReplacer(
	"\\\\", "\\",
	"\\`", "`",
	"\\*", "*",
	"\\_", "_",
	"\\{", "{",
	"\\}", "}",
	"\\[", "[",
	"\\]", "]",
	"\\<", "<",
	"\\>", ">",
	"\\(", "(",
	"\\)", ")",
	"\\#", "#",
	"\\+", "+",
	"\\-", "-",
	"\\.", ".",
	"\\!", "!",
	"\\|", "|",
)
