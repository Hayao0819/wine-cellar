package set

import "github.com/spf13/cobra"

func Root()(*cobra.Command){
	cmd := cobra.Command{
		Use: "set",
		Short: "Wine環境を設定",
		Long: "指定されたWineEnvに環境変数を設定します",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return &cmd
}
