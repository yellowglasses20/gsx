package gsx

import (
	"context"
	"io"
)

var (
	subCommands = []runner{
		&cmdPageList{},
		&cmdPageTitle{},
		&cmdPageText{},
		&cmdPageIcon{},
	}
	dispatch          = make(map[string]runner, len(subCommands))
	maxSubcommandName int
)

const cmdName = "gsx"

type runner interface {
	name() string
	description() string
	run(context.Context, []string, io.Writer, io.Writer) error
}

func init() {
	for _, r := range subCommands {
		n := r.name()
		l := len(n)
		if l > maxSubcommandName {
			maxSubcommandName = l
		}
		dispatch[n] = r
	}
}

// Run the gsx
func Run(argv []string, outStream, errStream io.Writer) error {

	return nil
}
