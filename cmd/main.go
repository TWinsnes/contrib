package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func Run() {
	app := &cli.Command{
		Name:    "contrib ",
		Usage:   "Generate a new Go CLI app scaffold powered by urfave/cli",
		Version: "0.1.0",
		// This add automated shell completion to the cli, see https://cli.urfave.org/v3/examples/completions/shell-completions/ for details
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			newCmd(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newCmd() *cli.Command {
	return &cli.Command{
		Name:      "hello",
		Usage:     "Greet the world",
		Action: func(c context.Context, cmd *cli.Command) error {
			fmt.Println("Hello, world!")
			return nil
		},
	}
}
