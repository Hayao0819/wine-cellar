package cmd

import (
	"os"
	"path"

	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/spf13/cobra"
	"github.com/Hayao0819/wine-cellar/cmd/env"
	"github.com/Hayao0819/wine-cellar/cmd/version"
		"github.com/Hayao0819/wine-cellar/cmd/run"
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
	}

	cmd.PersistentFlags().StringVarP(&conf.ConfDir, "config", "", func()(string){
		h, _:=os.UserHomeDir()
		return path.Join(h, ".wine-cellar")
	}() ,"config dir")

	cmd.AddCommand(
		env.Root(), 
		version.Root(),
		run.Root(),
	)

	return &cmd
}


func Execute() {
	err := rootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}
