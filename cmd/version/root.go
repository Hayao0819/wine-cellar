package version

import (
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "ver",
	Short: "Wine実行ファイルを管理",
	Long: `wineコマンドの追加や削除を行います。`,
}

