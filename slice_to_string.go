package gointerview

import (
	"fmt"
	"strings"
)

func sliceToBuilder[T fmt.Stringer](
	builder *strings.Builder, slice []T, sep string,
) {
	lastIndex := len(slice) - 1

	for i := 0; i < lastIndex; i++ {
		builder.WriteString(slice[i].String())
		builder.WriteString(sep)
	}

	builder.WriteString(slice[lastIndex].String())
}

func sliceToString[T fmt.Stringer](slice []T, sep string) string {
	var builder strings.Builder
	sliceToBuilder(&builder, slice, sep)
	return builder.String()
}
