// Copyright © 2024 Yuan Tong yt.vertigo0927@gmail.com

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	server bool
)

// portalCmd represents the portal command
var portalCmd = &cobra.Command{
	Use:   "portal",
	Short: "Portal CLI Util",
	Long:  `util 是一个用于Srun Portal CLI的快捷命令工具箱`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(server)
	},
}

func init() {
	portalCmd.PersistentFlags().BoolVarP(&server, "server", "s", false, "portal server")
	rootCmd.AddCommand(portalCmd)
}
