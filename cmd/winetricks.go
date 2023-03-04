package cmd

import (
	"os"
	"os/exec"

	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/spf13/cobra"
)


func runWineTricks(args ...string)(error){
	currentEnv, err := conf.GetCurrentEnv()
	if err !=nil{
		return err
	}
	
	winecmd := exec.Command("winetricks", args...)
	winecmd.Env=os.Environ()
	winecmd.Env=append(winecmd.Env, "WINEPREFIX="+currentEnv.Prefix)
	winecmd.Env=append(winecmd.Env, "WINEARCH="+currentEnv.Arch.Name)
	winecmd.Stdin=os.Stdin
	winecmd.Stderr=os.Stderr
	winecmd.Stdout=os.Stdout
	return winecmd.Run()
}

func newWineTricksCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use: "winetricks",
		Short: "Winetricksを実行します",
		Long: "現在のenvironment内でWinetricksを実行します",
		Aliases: []string{"tricks"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runWineTricks(args...)
		},
	}

	return &cmd
}
