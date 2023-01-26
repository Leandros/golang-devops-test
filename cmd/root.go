package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Verbose is set via CLI to enable verbose logging.
var Verbose bool
var Debug bool

// NewRootCommand creates the root command for Cobra CLI.
func NewRootCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands.
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "DevOps Test CLI",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(newDeployCmd())

	return rootCmd
}

// Execute executes the root command.
func Execute() {
	rootCmd := NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
