package command

import (
	"github.com/spf13/cobra"
)

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// app
	root.AddCommand(initAppCommand())
}
