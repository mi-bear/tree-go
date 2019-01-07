package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version  string
	Revision string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information of tree.",
	Run:   runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("version:    %s\n", Version)
	fmt.Printf("revision:   %s\n", Revision)
	fmt.Printf("go version: %s\n", runtime.Version())
	fmt.Printf("os/arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
