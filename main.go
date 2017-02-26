package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	outputType string
	debug      bool
)

var commands = []*Command{
	consume,
	publish,
}

// A Command is an implementation of a go command
// like go build or go fix.
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string) error

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'dnsme help' output.
	Short string

	// Long is the long message shown in the 'dnsme help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own
	// flag parsing.
	CustomFlags func(cmd *flag.FlagSet)
}

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}

func (c *Command) Runnable() bool {
	return c.Run != nil
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {

		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			addGlobalFlags(&cmd.Flag)
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags != nil {
				cmd.CustomFlags(&cmd.Flag)
			}
			//				args = args[1:]
			//			} else {
			cmd.Flag.Parse(args[1:])
			args = cmd.Flag.Args()
			//			}
			err := cmd.Run(cmd, args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown command %#q\n\n", args[0])
}

func addGlobalFlags(fs *flag.FlagSet) {
	fs.StringVar(&outputType, "o", "std", "Output type (std, json, csv)")
	fs.BoolVar(&debug, "d", false, "Debug output")
}
