package env

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Hayao0819/wine-cellar/conf"
)


func newListCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string)(error) {
			envs, err := conf.GetEnvs()
			if err !=nil{
				return err
			}
			
			for _, e:= range *envs{
				fmt.Printf("%s: WINEARCH=%s WINEPREFIX=%s CMD=%s\n", e.Name, e.Arch.Name, e.Prefix, e.Version.Cmd)
			}
			return nil
		},
	}

	return &cmd
}
