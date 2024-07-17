// Copyright © 2024 Yuan Tong <yt.vertigo0927@gmail.com>

package cmd

import (
	"github.com/dot-xiaoyuan/util/pkg/ssh"
	"github.com/spf13/cobra"
	"log"
)

var terminal ssh.Terminal

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use: "remote",
	Run: func(cmd *cobra.Command, args []string) {
		if terminal.Host == "" {
			log.Fatalf("host is required")
		}
		if err := terminal.New(); err != nil {
			log.Fatal(err)
		}
		terminal.Run("redis-cli -p 16384")
	},
}

func init() {

	rootCmd.AddCommand(remoteCmd)

	remoteCmd.PersistentFlags().StringVarP(&terminal.Host, "host", "h", "", "remote host")
	remoteCmd.PersistentFlags().Int32VarP(&terminal.Port, "port", "P", 22, "remote port")
	remoteCmd.PersistentFlags().StringVarP(&terminal.User, "username", "u", "root", "remote username")
	remoteCmd.PersistentFlags().StringVarP(&terminal.Password, "password", "p", "srunsoft@xian", "remote password")
}
