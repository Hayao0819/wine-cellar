package run

import "github.com/spf13/cobra"

func Root()(*cobra.Command){
	cmd := cobra.Command{
		Use: "run",
		Short: "コマンドを実行",
		Long: "指定されたツールを実行します",
	}

	for _, c := range []string{"winetricks", "notepad", "winecfg", "cmd",}{
		cmd.AddCommand(newCmd(c))
	}

	return &cmd
}
