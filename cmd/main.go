package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/TWinsnes/contrib/internal/contrib"
	"github.com/urfave/cli/v3"
)

func Run() {
	app := &cli.Command{
		Name:    "contrib ",
		Usage:   "Generate a contribution list for a git repository and write it to a file",
		Version: "0.1.0",
		// This adds automated shell completion to the cli, see https://cli.urfave.org/v3/examples/completions/shell-completions/ for details
		EnableShellCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Usage:   "The path to the file to write the contribution list to",
				Aliases: []string{"p"},
				Value:   "",
			},
		},
		Action: action,
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func action(ctx context.Context, cmd *cli.Command) error {
	c := contrib.NewContribOptions()

	path := cmd.String("path")

	if path != "" {
		c.Path = path
	}

	return c.Run(ctx)
}
