package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	pl := fmt.Println
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		pl("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		pl("subcommand 'foo'")
		pl("  enable:", *fooEnable)
		pl("  name:", *fooName)
		pl("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		pl("subcommand 'bar'")
		pl("  level:", *barLevel)
		pl("  tail:", barCmd.Args())
	default:
		pl("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
