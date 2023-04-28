package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/samber/lo"

	"go.szostok/see/internal"
)

func main() {

	cmd := flag.String("in", "", "")
	flag.Parse()
	raw := lo.Must(os.ReadFile("ex"))

	var examples []internal.Example

	lo.Must0(json.Unmarshal(raw, &examples))
	//width, _ := lo.Must2(term.GetSize(int(os.Stdout.Fd())))
	//
	//width /= 2
	//fmt.Println(width)
	for _, e := range examples {
		if !strings.EqualFold(*cmd, e.Command) {
			continue
		}

		r := lo.Must(glamour.NewTermRenderer(
			glamour.WithStylePath(getEnvironmentStyle()),
			glamour.WithWordWrap(130),
		))

		out := lo.Must(r.Render(e.Content))
		fmt.Print(out)
		break
	}
}

func getEnvironmentStyle() string {
	glamourStyle := os.Getenv("GLAMOUR_STYLE")
	if len(glamourStyle) == 0 {
		glamourStyle = "auto"
	}

	return glamourStyle
}
