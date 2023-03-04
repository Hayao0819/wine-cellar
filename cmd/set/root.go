package set

import (
	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/spf13/cobra"
)

func Root()(*cobra.Command){
	cmd := cobra.Command{
		Use: "set",
		Short: "Wine環境を設定",
		Long: "指定されたWineEnvをセットします",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			envs, err := conf.GetEnvs()
			if err !=nil{
				return err
			}

			e, err := envs.GetFromName(name)
			if err !=nil{
				return err
			}

			return conf.SetCurrentEnv(e)
		},
	}
	return &cmd
}
