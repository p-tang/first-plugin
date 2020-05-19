package main

import (
	"os"

	"github.com/p-tang/first-plugin/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	root := cmd.NewCmdVersion(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
