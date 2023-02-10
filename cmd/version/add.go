/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// env/addCmd represents the env/add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Wineコマンドを追加",
	Long: `WINEPREFIXをwine-cellarに登録し、管理対象にします。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("env/add called")
	},
}

func init() {
	VersionCmd.AddCommand(addCmd)
}
