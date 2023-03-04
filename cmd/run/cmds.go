package run

import (
	"os"
	"os/exec"

	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/spf13/cobra"
)

func newCmd(runcmd string)(*cobra.Command){
	return &cobra.Command{
		Use: runcmd,
		Short: "Run " + runcmd ,
		Long: "現在のWine環境内で"+runcmd+"を実行します",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentEnv, err := conf.GetCurrentEnv()
			if err !=nil{
				return err
			}

			println(currentEnv.Name)

			winecmd := exec.Command(currentEnv.Version.Cmd, append([]string{runcmd}, args...)...)
			winecmd.Env=os.Environ()
			winecmd.Env=append(winecmd.Env, "WINEPREFIX="+currentEnv.Prefix)
			winecmd.Env=append(winecmd.Env, "WINEARCH="+currentEnv.Arch.Name)
			winecmd.Stdin=os.Stdin
			winecmd.Stderr=os.Stderr
			winecmd.Stdout=os.Stdout
			return winecmd.Run()
		},
	}
}
