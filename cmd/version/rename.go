package version

import (
	"github.com/spf13/cobra"
)

func newRenameCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use: "rename OLD NEW",
		Short: "名前を変更する",
		Long: "Wineバージョンの名前を変更します",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	return &cmd
}
