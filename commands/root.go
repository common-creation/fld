package commands

import (
	"github.com/common-creation/fld/utils"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   utils.MustGetName(),
		Short: "Fast LSC Deployer",
	}
	cmd.AddCommand(NewDeployCmd())
	return cmd
}
