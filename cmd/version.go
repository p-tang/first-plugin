package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/projectcalico/calicoctl/calicoctl/commands"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func NewCmdVersion(streams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command {
		Use:		"kubectl calico version"
		Short:		"View calico version"
		Example:	"Client Version:    v3.14.0\n
					Git commit:        c97876ba\n
					Cluster Version:   v3.13.0\n
					Cluster Type:      k8s,canal,kdd\n"
		RunE:		func(c *cobra.Command, args []string) error {
			fmt.Println("Run Version Command")
		}
	}

	return cmd
}



