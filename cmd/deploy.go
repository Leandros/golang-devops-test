package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type newDeployOptions struct {
	to string
}

func newDeployCmd() *cobra.Command {
	options := newDeployOptions{
		to: "",
	}

	cmd := &cobra.Command{
		Use:     "deploy",
		Short:   "Deploy a project to an environment",
		Example: "deploy --to=staging frontend",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(1)(cmd, args); err != nil {
				return err
			}

			proj := args[0]
			if proj != "frontend" && proj != "backend" {
				return fmt.Errorf("project must be one of: [frontend, backend]")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if options.to != "dev" && options.to != "staging" && options.to != "prod" {
				return fmt.Errorf("to must be one of: [dev, staging, prod]")
			}

			fmt.Println("[=  ] deploying")

			time.Sleep(2000)

			fmt.Println("✔ deployed successfully!")

			return nil
		},
	}

	cmd.Flags().StringVar(&options.to, "to", "", "Which environment to deploy to")
	cmd.MarkFlagRequired("to")

	return cmd
}
