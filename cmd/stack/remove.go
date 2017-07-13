package stack

import (
	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
	"github.com/subfuzion/stack/stack"
)

func newRemoveCommand(dockerCli command.Cli) *cobra.Command {
	var opts stack.RemoveOptions

	cmd := &cobra.Command{
		Use:     "rm STACK [STACK...]",
		Aliases: []string{"remove", "down"},
		Short:   "Remove one or more stacks",
		Args:    cli.RequiresMinArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Namespaces = args
			return runRemove(dockerCli, opts)
		},
	}
	return cmd
}

func runRemove(dockerCli command.Cli, opts stack.RemoveOptions) error {
	return stack.Remove(dockerCli, opts)
}
