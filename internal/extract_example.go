package internal

import (
	"bytes"
	"strings"
)

func ProcessSimpleExamples(in string) []Example {
	lines := strings.FieldsFunc(in, splitByNewLines)

	var out []Example
	var buff bytes.Buffer
	for idx := range lines {
		l := lines[idx]
		l = strings.TrimSpace(l)

		if strings.HasPrefix(l, "#") {
			l = strings.TrimPrefix(l, "#")
			l = strings.TrimSpace(l)
			buff.WriteString(" " + l)
			continue
		}

		out = append(out, Example{
			Command:     l,
			Description: buff.String(),
		})
		buff.Reset()
	}

	return out
}

func splitByNewLines(c rune) bool {
	return c == '\n' || c == '\r'
	//unicode.Zl
	//return !unicode.IsOneOf(c) && !unicode.IsNumber(c)
}
