package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tree-go",
	Short: "files are listed.",
	Args:  cobra.ArbitraryArgs,
	Run:   runRoot,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

func runRoot(cmd *cobra.Command, args []string) {
	path, err := getFilePath(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", path)

	dirs := make([]string, 0)

	if err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if path == p {
				return nil
			}

			list := strings.Split(p, "/")
			node := len(list)

			if node > 1 {
				fmt.Printf("│")
			}

			x := node
			for x > 1 {
				x--
				fmt.Printf("   ")
			}

			fmt.Printf("├── ")
			fmt.Println(list[node-1])

			dirs = append(dirs, list[node-1])
			return nil
		}

		rel, err := filepath.Rel(path, p)
		if err != nil {
			return err
		}

		list := strings.Split(rel, "/")
		node := len(list)

		if node > 1 {
			fmt.Printf("│")
		}

		x := node
		for x > 1 {
			x--
			fmt.Printf("   ")
		}

		fmt.Printf("├── ")
		fmt.Println(list[node-1])

		return nil
	}); err != nil {
		fmt.Println(err)
	}
}

func getFilePath(args []string) (string, error) {
	if len(args) < 1 {
		return ".", nil
	}

	arg := strings.TrimSpace(args[0])
	if arg == "." {
		return arg, nil
	}

	p, err := filepath.Abs(filepath.Clean(arg))
	if err != nil {
		return "", err
	}

	_, err = os.Stat(p)
	if err != nil {
		return "", err
	}

	return p, nil
}
