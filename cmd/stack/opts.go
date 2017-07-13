package stack

import (
	"github.com/spf13/pflag"
)

func addComposefileFlag(opt *string, flags *pflag.FlagSet) {
	flags.StringVarP(opt, "compose-file", "c", "", "Path to a Compose file")
	flags.SetAnnotation("compose-file", "version", []string{"1.25"})
}

func addBundlefileFlag(opt *string, flags *pflag.FlagSet) {
	flags.StringVar(opt, "bundle-file", "", "Path to a Distributed Application Bundle file")
	flags.SetAnnotation("bundle-file", "experimental", nil)
}

func addRegistryAuthFlag(opt *bool, flags *pflag.FlagSet) {
	flags.BoolVar(opt, "with-registry-auth", false, "Send registry authentication details to Swarm agents")
}

