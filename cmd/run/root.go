package run

import "github.com/spf13/cobra"

func Root()(*cobra.Command){
	cmd := cobra.Command{
		Use: "run",
		Short: "コマンドを実行",
		Long: "指定されたツールを実行します",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runWine(args[0], args[1:]...)
		},
	}

	/*
	for _, c := range []string{"winetricks", "notepad", "winecfg", "cmd",}{
		cmd.AddCommand(newCmd(c))
	}
	*/

	return &cmd
}
