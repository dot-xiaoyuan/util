/*
Copyright © 2024 Yuan Tong <yt.vertigo0927@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var host string
var port int32
var remoteUser string
var remotePassword string

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Srun4k redis-cli",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		cliCmd := exec.Command("which", "redis-cli1")
		output, err := cliCmd.CombinedOutput()
		if err != nil {
			// localhost not have redis-cli
			// try to ssh remote server
		}
		cliPath := strings.TrimSpace(string(output))

		c := exec.Command(cliPath, "-h", host, "-p", strconv.Itoa(int(port)))
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			fmt.Println("Error executing command:", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("help", "", false, "help for this command")
	rootCmd.AddCommand(redisCmd)

	redisCmd.PersistentFlags().StringVarP(&host, "host", "h", "127.0.0.1", "redis server host")
	redisCmd.PersistentFlags().Int32VarP(&port, "port", "P", 6379, "redis server port")
	redisCmd.PersistentFlags().StringVarP(&remoteUser, "user", "u", "root", "remote server ssh user")
	redisCmd.PersistentFlags().StringVarP(&remotePassword, "password", "p", "srunsoft@xian", "remote server ssh password")

}
