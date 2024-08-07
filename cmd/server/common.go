package server

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	Portal         = "portal"
	PortalSvr      = "portalsvr"
	Srun3kAuthCtrl = "/srun3/bin/srun3kauth-ctrl.sh"
)

func ServiceCmd(u, s string, runFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:   u,
		Short: fmt.Sprintf("%s %s server", u, s),
		Run:   runFunc,
	}
}
