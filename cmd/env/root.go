package env

import (
	"github.com/spf13/cobra"
)



func Root ()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "env",
		Short: "WINEPREFIXを管理",
		Long: `WINEPREFIXの追加や削除を行います。`,
	}

	cmd.AddCommand(newListCmd(), newAddCmd())

	return &cmd
}
