package template

import (
	"github.com/spf13/cobra"
)

const usage = `{{h1 "Usage"}}{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

{{h1 "Aliases"}}
{{.NameAndAliases}}{{end}}{{if .HasExample}}

{{h1 "Examples"}}
{{removeCode .Example}}{{end}}{{if .HasAvailableSubCommands}}

{{h1 "Available Commands"}}{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

{{h1 "Flags"}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

{{h1 "Global Flags"}}
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

// Usage is the text template for the root command.
func Usage() string {
	cobra.AddTemplateFunc("h1", h1)
	cobra.AddTemplateFunc("removeCode", removeCode)
	cobra.AddTemplateFunc("rpad", rpad)
	return usage
}
