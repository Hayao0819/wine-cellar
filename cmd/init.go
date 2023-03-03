package cmd

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/Hayao0819/wine-cellar/cmd/version"
	"github.com/spf13/cobra"
)

func newInitCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use: "init",
		Short: "初期化",
		Long: "空のファイルを作成して初期化します",
		RunE: func(cmd *cobra.Command, args []string) error {
			errs := []string{}
			envConfDir := path.Dir(conf.EnvConfFile)
			verConfDir := path.Dir(conf.EnvConfFile)

			if envConfDir == conf.ConfDir && verConfDir == conf.ConfDir{
				if err :=os.Mkdir(conf.ConfDir, 0750); err !=nil{
					errs=append(errs, err.Error()) 
				}
			}else if envConfDir == verConfDir{
				if err :=os.Mkdir(envConfDir, 0750); err !=nil{
					errs=append(errs, err.Error()) 
				}
			}else{
				if err :=os.Mkdir(envConfDir, 0750); err !=nil{
					errs=append(errs, err.Error()) 
				}

				if err:= os.Mkdir(verConfDir, 0750); err!=nil{
					errs=append(errs, err.Error()) 
				}
			}


			env, err := os.Create(conf.EnvConfFile)
			if err !=nil{
				errs=append(errs, err.Error()) 
			}
			defer env.Close()

			ver, err := os.Create(conf.VerConfFile)
			if err !=nil{
				errs=append(errs, err.Error()) 
			}
			defer ver.Close()



			versionCmd := version.Root()
			versionCmd.SetArgs([]string{"update"})
			if err:= versionCmd.Execute(); err !=nil{
				errs = append(errs, err.Error())
			}

			if len(errs)>0{
				return errors.New(strings.Join(errs, "\n"))
			}
			return nil
		},
	}
	return &cmd
}
