package stack

import (
	"fmt"
	"io"
	"os"

	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/bundlefile"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/subfuzion/stack/stack"
)

func newDeployCommand(dockerCli command.Cli) *cobra.Command {
	var opts stack.DeployOptions

	cmd := &cobra.Command{
		Use:     "deploy [OPTIONS] STACK",
		Aliases: []string{"up"},
		Short:   "Deploy a new stack or update an existing stack",
		Args:    cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Namespace = args[0]
			return runDeploy(dockerCli, opts)
		},
	}

	flags := cmd.Flags()
	addBundlefileFlag(&opts.Bundlefile, flags)
	addComposefileFlag(&opts.Composefile, flags)
	addRegistryAuthFlag(&opts.SendRegistryAuth, flags)
	flags.BoolVar(&opts.Prune, "prune", false, "Prune services that are no longer referenced")
	flags.SetAnnotation("prune", "version", []string{"1.27"})
	flags.StringVar(&opts.ResolveImage, "resolve-image", stack.ResolveImageAlways,
		`Query the registry to resolve image digest and supported platforms ("`+stack.ResolveImageAlways+`"|"`+stack.ResolveImageChanged+`"|"`+stack.ResolveImageNever+`")`)
	flags.SetAnnotation("resolve-image", "version", []string{"1.30"})
	return cmd
}

func runDeploy(dockerCli command.Cli, opts stack.DeployOptions) error {
	return stack.Deploy(dockerCli, opts)
}

func loadBundlefile(stderr io.Writer, namespace string, path string) (*bundlefile.Bundlefile, error) {
	defaultPath := fmt.Sprintf("%s.dab", namespace)

	if path == "" {
		path = defaultPath
	}
	if _, err := os.Stat(path); err != nil {
		return nil, errors.Errorf(
			"Bundle %s not found. Specify the path with --file",
			path)
	}

	fmt.Fprintf(stderr, "Loading bundle from %s\n", path)
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	bundle, err := bundlefile.LoadFile(reader)
	if err != nil {
		return nil, errors.Errorf("Error reading %s: %v\n", path, err)
	}
	return bundle, err
}
