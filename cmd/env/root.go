package env

import (
	"github.com/spf13/cobra"
)

var EnvCmd = &cobra.Command{
	Use:   "env",
	Short: "WINEPREFIXを管理",
	Long: `WINEPREFIXの追加や削除を行います。`,
}
