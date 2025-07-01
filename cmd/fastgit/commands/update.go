package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/update"
	"github.com/spf13/cobra"
)

func UpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update to latest version",
		Run: func(cmd *cobra.Command, args []string) {
			if err := update.UpdateSelf(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
