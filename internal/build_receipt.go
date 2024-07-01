package internal

import "strings"

// Writes all receipt lines from col to builder
func WriteCollection(builder *strings.Builder, col *IteratorCollection) bool {
	if col.ok {
		builder.WriteString("(OK) ")
	} else {
		builder.WriteString("(  ) ")
	}

	center := (col.maxHeight - (1 - (col.maxHeight & 1))) / 2
	last := col.maxHeight - 1

	for i := 0; i < col.maxHeight; i++ {
		if i > 0 {
			builder.WriteString("     ")
		}

		col.WriteInput(builder, i)

		if col.input2 != nil {
			if i == center {
				builder.WriteString(", ")
			} else {
				builder.WriteString("  ")
			}
			col.WriteInput2(builder, i)
		}

		if i == center {
			builder.WriteString(" -> ")
		} else {
			builder.WriteString("    ")
		}

		col.WriteActual(builder, i)

		if !col.ok {
			if i == center {
				builder.WriteString(" != ")
			} else {
				builder.WriteString("    ")
			}

			col.WriteExpected(builder, i)
		}

		if i < last {
			builder.WriteRune('\n')
		}
	}

	return col.maxHeight > 1
}
