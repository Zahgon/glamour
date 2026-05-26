package ansi

import (
	"github.com/yuin/goldmark/ast"
)

type tableLink struct {
	href     string
	title    string
	content  string
	linkType linkType
}

type linkType int

const (
	_ linkType = iota
	linkTypeAuto
	linkTypeImage
	linkTypeRegular
)

func (e *TableElement) printTableLinks(ctx RenderContext) { _ = "STUB: not implemented"; return }

//nolint: gosec

func (e *TableElement) shouldPrintTableLinks(ctx RenderContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *TableElement) collectLinksAndImages(ctx RenderContext) error {
	_ = "STUB: not implemented"
	return nil
}

func isInsideTable(node ast.Node) bool { _ = "STUB: not implemented"; return false }

func nodeContent(node ast.Node, source []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func linkDomain(href string) string { _ = "STUB: not implemented"; return "" }

func linkWithSuffix(tl tableLink, list []tableLink) string { _ = "STUB: not implemented"; return "" }
