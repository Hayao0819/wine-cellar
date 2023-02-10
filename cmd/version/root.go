package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "env",
	Short: "Wine実行ファイルを管理",
	Long: `wineコマンドの追加や削除を行います。`,
}

func Execute() {
	err := VersionCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}