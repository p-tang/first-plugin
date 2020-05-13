package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var version string

func SetVersion(v string) {
	version = v
}

type versionCmd struct {
	out io.Writer
}

func newVersionCmd(out io.Writer) *cobra.Command {
	version := &versionCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "print the version number and then exit",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return errors.New("This command does not take arguments")
			}
			return version.run()
		},
	}
	return cmd
}

func (v *versionCmd) run() error {
	_, err := fmt.Fprintf(v.out, "kubectl server-version %s\n", version)
	if err != nil {
		return err
	}

	return nil
}
