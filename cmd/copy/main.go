package main

import (
	"encoding/json"
	"fmt"
	"os"

	fzf "github.com/junegunn/fzf/src"
	"github.com/junegunn/fzf/src/protector"
	"github.com/samber/lo"

	"go.szostok/see/internal"
)

var version string = "0.34"
var revision string = "devel"

func main() {
	protector.Protect()
	os.Setenv("FZF_DEFAULT_OPTS", `--multi --reverse --preview="cat ./ex  | jq -r '.[] | select(.Command==\"{}\") | .Content' | bat -l md -f -n -H 19" --preview-window=right:60%,wrap -i`)

	raw := lo.Must(os.ReadFile("ex"))

	var examples []internal.Example

	lo.Must0(json.Unmarshal(raw, &examples))
	for _, e := range examples {
		os.Stdin.WriteString(fmt.Sprintln(e.Command))
	}
	fzf.Run(fzf.ParseOptions(), version, revision)
}
