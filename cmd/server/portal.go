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

package server

import (
	"github.com/dot-xiaoyuan/util/pkg/util"
	"github.com/spf13/cobra"
	"os/exec"
)

// PortalCmd represents the server/portal command
var PortalCmd = &cobra.Command{
	Use:   "portal",
	Short: "portal server util",
	Long:  `适用于Srun Portal的CLI工具`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func init() {
	PortalCmd.AddCommand(ServiceCmd("start", Portal, Start))
	PortalCmd.AddCommand(ServiceCmd("stop", Portal, Stop))
	PortalCmd.AddCommand(ServiceCmd("status", Portal, Status))
	PortalCmd.AddCommand(ServiceCmd("restart", Portal, Restart))
}

func Start(cmd *cobra.Command, args []string) {
	c := exec.Command(Srun3kAuthCtrl, "start-"+PortalSvr)
	util.ShowCommandOutputWithSpinner(c, "Starting the server...")
}

func Stop(cmd *cobra.Command, args []string) {
	c := exec.Command(Srun3kAuthCtrl, "stop-"+PortalSvr)
	util.ShowCommandOutput(c)
}

func Status(cmd *cobra.Command, args []string) {
	c := exec.Command(Srun3kAuthCtrl, "status")
	util.ShowCommandOutput(c)
}

func Restart(cmd *cobra.Command, args []string) {
	c := exec.Command(Srun3kAuthCtrl, "restart-"+PortalSvr)
	util.ShowCommandOutputWithSpinner(c, "Restarting the server...")
}
