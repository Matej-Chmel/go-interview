package internal

import "strings"

// Iterates lines from a result that.
// Represents an input or actual or expected output.
type LineIterator struct {
	current    int
	lines      []string
	height     int
	skipString string
	startAt    int
	width      int
}

// Calculates the first line that is not empty from mh,
// maximum height of all iterators.
func (i *LineIterator) calculateSkip(mh int) {
	i.startAt = (mh - i.height) / 2
	i.skipString = strings.Repeat(" ", i.width)
}

// Returns maximum width from lines.
func maxWidth(lines []string) (w int) {
	for _, l := range lines {
		w = max(w, len(l))
	}

	return
}

// Constructs a new LineIterator
func NewLinesIterator(s string) *LineIterator {
	lines := strings.Split(s, "\n")

	return &LineIterator{
		current:    0,
		lines:      lines,
		height:     len(lines),
		skipString: "",
		startAt:    -1,
		width:      maxWidth(lines),
	}
}

// Returns the next available line
func (it *LineIterator) Next(i int) string {
	if i < it.startAt {
		return it.skipString
	}

	res := it.lines[it.current]
	it.current++
	return res
}

// Collection of iterators for inputs and outputs
type IteratorCollection struct {
	actual    *LineIterator
	expected  *LineIterator
	input     *LineIterator
	input2    *LineIterator
	ok        bool
	maxHeight int
}

// Constructs a new IteratorCollection
func NewIteratorCollection(actual, expected string, input1 string, input2 *string) *IteratorCollection {
	var i2 *LineIterator = nil

	if input2 != nil {
		i2 = NewLinesIterator(*input2)
	}

	c := &IteratorCollection{
		actual:    NewLinesIterator(actual),
		expected:  NewLinesIterator(expected),
		input:     NewLinesIterator(input1),
		input2:    i2,
		ok:        actual == expected,
		maxHeight: 0,
	}
	c.calculateSkip()
	return c
}

// Calculates the first non-empty line for all iterators
func (c *IteratorCollection) calculateSkip() {
	c.maxHeight = max(c.actual.height, c.expected.height, c.input.height)

	if c.input2 != nil {
		c.maxHeight = max(c.maxHeight, c.input2.height)
		c.input2.calculateSkip(c.maxHeight)
	}

	c.actual.calculateSkip(c.maxHeight)
	c.expected.calculateSkip(c.maxHeight)
	c.input.calculateSkip(c.maxHeight)
}

// Writes next available line from iterator for the actual output to b
func (c *IteratorCollection) WriteActual(b *strings.Builder, i int) {
	c.writeString(b, c.actual, i, IsActual)
}

// Writes next available line from iterator for the expected output to b
func (c *IteratorCollection) WriteExpected(b *strings.Builder, i int) {
	c.writeString(b, c.expected, i, IsExpected)
}

// Writes next available line from iterator for first input to b
func (c *IteratorCollection) WriteInput(b *strings.Builder, i int) {
	c.writeString(b, c.input, i, 0)
}

// Writes next available line from iterator for second input to b
func (c *IteratorCollection) WriteInput2(b *strings.Builder, i int) {
	c.writeString(b, c.input2, i, 0)
}

// Write a line and its right padding unless the string is the last one
// on the current line in the string builder b
func (c *IteratorCollection) writeString(
	b *strings.Builder, iter *LineIterator, i int, flags uint,
) {
	if iter == nil {
		return
	}

	data := iter.Next(i)
	b.WriteString(data)

	if len(data) >= iter.width {
		return
	}

	isActualLast := c.ok && (flags&IsActual) == IsActual
	isExpectedLast := !c.ok && (flags&IsExpected) == IsExpected

	if isActualLast || isExpectedLast {
		return
	}

	diff := iter.width - len(data)

	for j := 0; j < diff; j++ {
		b.WriteRune(' ')
	}
}
