package cmd

import (
	"os"
	"path"
	//"fmt"

	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/spf13/cobra"
	"github.com/Hayao0819/wine-cellar/cmd/env"
	"github.com/Hayao0819/wine-cellar/cmd/version"
	"github.com/Hayao0819/wine-cellar/cmd/run"
	"github.com/Hayao0819/wine-cellar/cmd/set"
)

func rootCmd ()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "wine-cellar",
		Short: "Manage your wine environment",
		Long: `Wine-cellarはGo言語で書かれたWINEPREFIXの管理ツールです。`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//Run: func(cmd *cobra.Command, args []string) { },
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			conf.Initilize()
		},
		/*
		RunE: func(cmd *cobra.Command, args []string) error {
			e, err := conf.GetCurrentEnv()
			if err !=nil{
				return err
			}
			fmt.Println(e.Name)
			return nil
		},
		*/
	}

	cmd.PersistentFlags().StringVarP(&conf.ConfDir, "config", "", func()(string){
		h, _:=os.UserHomeDir()
		return path.Join(h, ".wine-cellar")
	}() ,"config dir")

	cmd.AddCommand(
		env.Root(), 
		version.Root(),
		run.Root(),
		set.Root(),
		newInitCmd(),
		newWineTricksCmd(),
	)

	return &cmd
}


func Execute() {
	err := rootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}
