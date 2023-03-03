package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newAddCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "add",
		Short: "Wineコマンドを追加",
		Long: `WINEPREFIXをwine-cellarに登録し、管理対象にします。`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("env/add called")
		},
	}
	return &cmd
}
