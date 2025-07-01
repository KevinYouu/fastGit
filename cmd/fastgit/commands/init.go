package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/spf13/cobra"
)

func InitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize configuration",
		Run: func(cmd *cobra.Command, args []string) {
			if err := config.Initialize(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
