package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var EnvCmd = &cobra.Command{
	Use:   "env",
	Short: "WINEPREFIXを管理",
	Long: `WINEPREFIXの追加や削除を行います。`,
}

func Execute() {
	err := EnvCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
