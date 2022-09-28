// Package template provides usage templates to render help menus.
package template

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func h1(text string) string {
	return color.New(color.Bold).Sprint(strings.ToUpper(text))
}

func removeCode(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = fmt.Sprintf("%s%s", strings.Repeat(" ", 2), strings.ReplaceAll(line, "/code ", ""))
	}
	return strings.Join(lines, "\n")
}

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}
