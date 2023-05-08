package version

import (
	"github.com/spf13/cobra"
)



func Root()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "ver",
		Short: "Wine実行ファイルを管理",
		Long: `wineコマンドの追加や削除を行います。`,
	}

	cmd.AddCommand(newAddCmd() ,newListCmd(), newUpdateCmd(), newRenameCmd(), newInstallCmd())

	return &cmd
}
