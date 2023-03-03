package version

import (
	"fmt"

	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/spf13/cobra"
)

func newListCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use: "list",
		Short: "バージョン一覧を表示",
		Long: "wine-cellarが認識しているバージョン一覧を表示します",
		RunE: func(cmd *cobra.Command, args []string) error {
			vers, err := conf.GetVersions()
			if err !=nil{
				return err
			}

			for _, v := range *vers{
				fmt.Println(v.Name)
			}
			return nil
		},
	}

	return &cmd
}
