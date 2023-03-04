package env

import (
	"errors"

	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/Hayao0819/wine-cellar/go-wine"

	"github.com/spf13/cobra"
)


func newAddCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "add NAME ARCH PATH VERSION",
		Short: "WINEPREFIXを登録",
		Long: `WINEPREFIXをwine-cellarに登録し、管理対象にします。`,
		Args: cobra.ExactArgs(4), // [name] [arch] [path] [version]
		RunE: func(cmd *cobra.Command, args []string) error{
			envs, err := conf.GetEnvs()
			if err !=nil{
				return err
			}

			ver, err := func()(*wine.Version, error){
				vers, err := conf.GetVersions()
				if err !=nil{
					return nil, err
				}
				for _,v := range *vers{
					if v.Name==args[3]{
						return &v, nil
					}
				}
				return nil, errors.New("cant find such version")
			}()
			if err!=nil{
				return err
			}


			new := wine.NewEnv(args[0], args[1], args[2], *ver)

			*envs=append(*envs, new) //具体的な情報付与はあとで実装する

			return conf.WriteEnvs(envs)
		},
	}

	return &cmd
}
