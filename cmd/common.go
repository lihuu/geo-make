package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"path"
)

func PersistentPreFunc(cmd *cobra.Command, args []string) error {
	if srcDir != "" {
		return nil
	}
	srcDir, err := os.Getwd()
	if err != nil {
		return err
	}
	_ = MakeDir(path.Join(srcDir, "ips"))
	_ = MakeDir(path.Join(srcDir, "sites"))
	return processOutDir()
}

func MakeDir(srcDir string) error {
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		err = os.Mkdir(srcDir, DirPerm)
		if err != nil {
			return err
		}
	}
	return nil
}
