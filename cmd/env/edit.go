package env

import (
	"strings"

	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/Hayao0819/wine-cellar/go-wine"
	"github.com/spf13/cobra"
)


func newEditCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use: "edit KEY VALUE",
		Short: "wine envを編集します",
		Long: "現在選択されているenvを編集します",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			envs , err := conf.GetEnvs()
			if err !=nil{
				return err
			}

			current, err := conf.GetCurrentEnv()
			if err !=nil{
				return err
			}

			key := strings.ToLower(args[0]) 
			value :=args[1]

			switch key{
				case "name":
					current.Name=value
				case "arch":
					switch value{
						case "win64": 
							current.Arch=wine.Win64
						case "win32":
							current.Arch=wine.Win32
						default:
							return wine.ErrNoSuchArch
					}
				case "prefix":
					current.Prefix=value
				case "version":
					// あとで実装
			}

			// 書き込みをあとで実装

			return conf.WriteEnvs(envs)
		},
	}

	return &cmd
}
