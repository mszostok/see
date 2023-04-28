package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Netflix/go-expect"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/samber/lo"
)

func main() {
	//
	//args := lo.Must(shellwords.Parse(`bash -c """`))
	cmdRaw := `cat ./ex  | jq '.[].Command' | fzf --multi --reverse --preview="cat ./ex  | jq -r '.[] | select(.Command==\"{}\") | .Content' | bat -l md -f -n -H 19" --preview-window=right:60%,wrap -i`
	encoding.Register()

	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	cmd := exec.Command("bash", "-c", cmdRaw)
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()
	lo.Must0(cmd.Start())

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				os.Exit(0)
			}
			ev.Modifiers()
			c.Send("\x001B")
		}
	}

	//cmd.Stdin = os.Stdin

	//fmt.Println(cmd.String())

}
