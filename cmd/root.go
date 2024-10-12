package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	srcDir string
	outDir string
)

var rootCmd = &cobra.Command{
	Use:   "geo-make",
	Short: "Geo Make is a tool to generate geo resources from plain text files",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
