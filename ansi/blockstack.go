package ansi

// BlockStack is a stack of block elements, used to calculate the current
// indentation & margin level during the rendering process.
type BlockStack []BlockElement

// Len returns the length of the stack.
func (s *BlockStack) Len() int {
	_ = "STUB: not implemented"

	// Push appends an item to the stack.
	return 0
}

func (s *BlockStack) Push(e BlockElement) {
	_ = "STUB: not implemented"

	// Pop removes the last item on the stack.
	return
}

func (s *BlockStack) Pop() { _ = "STUB: not implemented"; return }

// Indent returns the current indentation level of all elements in the stack.
func (s BlockStack) Indent() uint { _ = "STUB: not implemented"; return 0 }

// Margin returns the current margin level of all elements in the stack.
func (s BlockStack) Margin() uint { _ = "STUB: not implemented"; return 0 }

// Width returns the available rendering width.
func (s BlockStack) Width(ctx RenderContext) uint { _ = "STUB: not implemented"; return 0 }

//nolint: gosec

//nolint: gosec

// Parent returns the current BlockElement's parent.
func (s BlockStack) Parent() BlockElement { _ = "STUB: not implemented"; return *new(BlockElement) }

// Current returns the current BlockElement.
func (s BlockStack) Current() BlockElement { _ = "STUB: not implemented"; return *new(BlockElement) }

// With returns a StylePrimitive that inherits the current BlockElement's style.
func (s BlockStack) With(child StylePrimitive) StylePrimitive {
	_ = "STUB: not implemented"
	return *new(StylePrimitive)
}
