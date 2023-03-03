package version

import (
	"errors"

	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/Hayao0819/wine-cellar/go-wine"
	"github.com/spf13/cobra"
)

func newAddCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "add",
		Short: "Wineコマンドを追加",
		Long: `WINEPREFIXをwine-cellarに登録し、管理対象にします。`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string)error {
			vers, err := conf.GetVersions()
			if err !=nil{
				return err
			}

			new, err := wine.GetLocalVersion(args[0])
			if err !=nil{
				return err
			}

			// 同じ項目がないかチェック
			for _, v := range *vers{
				if v.Name==new.Name || v.Cmd==new.Cmd{
					return errors.New("duplicate name or command")
				}
			}

			*vers=append(*vers, *new)

			return conf.WriteVersions(vers)
		},
	}
	return &cmd
}
