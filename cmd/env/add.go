package env

import (
	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/Hayao0819/wine-cellar/go-wine"

	"github.com/spf13/cobra"
)


func newAddCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "add",
		Short: "WINEPREFIXを登録",
		Long: `WINEPREFIXをwine-cellarに登録し、管理対象にします。`,
		Args: cobra.ExactArgs(2), // [name] [path]
		RunE: func(cmd *cobra.Command, args []string) error{
			envs, err := conf.GetEnvs()
			if err !=nil{
				return err
			}

			*envs=append(*envs, wine.Env{}) //具体的な情報付与はあとで実装する

			return conf.WriteEnvs(envs)
		},
	}

	return &cmd
}
