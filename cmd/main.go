package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/mattn/go-shellwords"
	"github.com/muesli/reflow/wordwrap"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"k8s.io/kubectl/pkg/cmd"

	"go.szostok/see/internal"
	"go.szostok/see/internal/template"
)

var what = cobra.Command{
	Aliases:           nil,
	SuggestFor:        nil,
	Short:             "",
	Long:              "",
	Example:           "",
	ValidArgs:         nil,
	ValidArgsFunction: nil,
	ArgAliases:        nil,
	// SuggestionsMinimumDistance defines minimum levenshtein distance to display suggestions.
	// Must be > 0.
	//SuggestionsMinimumDistance: 0,
}

type Hoard struct {
	Version  string     `yaml:"version"`
	Commands []Commands `yaml:"commands"`
}
type Commands struct {
	Name        string   `yaml:"name"`
	Namespace   string   `yaml:"namespace"`
	Tags        []string `yaml:"tags"`
	Command     string   `yaml:"command"`
	Description string   `yaml:"description"`
}

func main() {
	db := internal.DB{}

	identity := func(s string) string { return s }
	emptyStr := func(s string) string { return "" }

	err := template.GenMarkdownTreeCustom(cmd.NewDefaultKubectlCommand(), &db, emptyStr, identity)
	//err := internal.GenMarkdownTreeCustom(cmd.NewDefaultKubectlCommand(), &db)
	if err != nil {
		log.Fatal(err)
	}

	// ccat ./ex  | jq '.[].Command' | fzf --multi --reverse --preview="cat ./ex  | jq -r '.[] | select(.Command==\"{}\") | .Content'"
	data, err := json.Marshal(db.Examples)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("ex", data, 0o655)
	if err != nil {
		log.Fatal(err)
	}

	return
	h := Hoard{}

	for _, item := range db.Examples {
		out, _ := glamour.Render(item.Content, "dark")
		h.Commands = append(h.Commands, Commands{
			Name:        strings.ReplaceAll(item.Command, " ", "-"),
			Namespace:   "kubectl",
			Tags:        []string{"kubectl"},
			Command:     item.Command,
			Description: out,
		})
	}

	h.Version = "1.0.1"
	data, err = yaml.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("examples", data, 0o655)
	if err != nil {
		log.Fatal(err)
	}

	return
	idx, err := fuzzyfinder.Find(
		db.Examples,
		func(i int) string {
			return fmt.Sprintf("%s [%s]", db.Examples[i].Description, db.Examples[i].Command)
		},
		fuzzyfinder.WithCursorPosition(fuzzyfinder.CursorPositionTop),
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}

			out, _ := glamour.Render(db.Examples[i].Content, "dark")
			return fmt.Sprintf("Description %s\nDescription %s\nCommand: %s",
				wordwrap.String(out, w/2-20),
				wordwrap.String(db.Examples[i].Description, w/2-20),
				wordwrap.String(db.Examples[i].Command, w),
			)
		}))
	if err != nil {
		log.Fatal(err)
	}

	rawCmd := db.Examples[idx].Command
	if rawCmd == "" {
		log.Fatal("no command to run")
	}
	envs, args, err := shellwords.ParseWithEnvs(rawCmd)

	cmd := exec.Command(args[0], args[1:]...)
	for _, env := range envs {
		cmd.Env = append(cmd.Env, env)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
