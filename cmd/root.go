package cmd

import (
	"fmt"
	"io/ioutil"
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

	if path == "." {
		cd, _ := os.Getwd()
		path = cd
	}

	dirCount := 0
	fileCount := 0

	dirs := make(map[string]int)
	files := make(map[string]int)

	if err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		parentDir := filepath.Dir(p)
		files[parentDir]++
		last := files[parentDir] == dirs[parentDir]

		if info.IsDir() {
			f, err := ioutil.ReadDir(p)
			if err != nil {
				return err
			}
			dirs[p] = len(f)

			if path == p {
				return nil
			}

			list := strings.Split(strings.Replace(p, path+"/", "", -1), "/")
			node := len(list)

			filePrint(list[node-1], node, last)

			dirCount++
			return nil
		}

		rel, err := filepath.Rel(path, p)
		if err != nil {
			return err
		}

		list := strings.Split(rel, "/")
		node := len(list)

		filePrint(list[node-1], node, last)

		fileCount++
		return nil
	}); err != nil {
		fmt.Println(err)
	}

	lastPrint(dirCount, fileCount)
}

func filePrint(file string, node int, last bool) {
	if node > 1 {
		fmt.Printf("│")
		fmt.Printf("   ")
	}

	x := node - 2
	for x > 0 {
		x--
		fmt.Printf("│")
		fmt.Printf("   ")
	}

	if last {
		fmt.Printf("└── ")
	} else {
		fmt.Printf("├── ")
	}
	fmt.Println(file)
}

func lastPrint(dirCount, fileCount int) {
	dir := "directory"
	file := "file"
	if dirCount != 1 {
		dir = "directories"
	}
	if fileCount != 1 {
		file = "files"
	}
	fmt.Printf("\n%d %s, %d %s", dirCount, dir, fileCount, file)
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
