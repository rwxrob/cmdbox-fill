package cmd

import (
	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/cmdbox/fmt"
)

func init() {
	x := cmdbox.New("fill")
	x.Summary = `fill simple text form templates`
	x.Usage = `[<field> ...]`
	x.Version = `v1.0.0`

	x.Description = `
		The *fill* command simply replaces all {n} fields within the text
		sent from standard input with the associated argument (starting with
		#1). Use care when passing text that might accidentally have such
		fields mistakenly replaced. `

	x.Method = func(args []string) (err error) {
		if len(args) == 0 {
			return x.UsageError()
		}
		fmt.Print(fmt.FillIn(args...))
		return nil
	}
}
