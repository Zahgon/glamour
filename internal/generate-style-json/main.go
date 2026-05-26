package main //nolint:revive

import (
	"fmt"
	"os"

	"charm.land/glamour/v2/ansi"
)

func writeStyleJSON(filename string, styleConfig *ansi.StyleConfig) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint: errcheck

func run() error { _ = "STUB: not implemented"; return nil }

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
