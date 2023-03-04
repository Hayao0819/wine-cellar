package version

import (
	"errors"
	"strings"

	"github.com/Hayao0819/wine-cellar/conf"
	"github.com/Hayao0819/wine-cellar/go-wine"
	"github.com/spf13/cobra"
)

func newUpdateCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use: "update",
		Short: "バージョン情報を更新します",
		Long: "システムの実際のバージョンと設定ファイルの情報を同期します",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 現在のバージョン一覧を取得
			vers, err := conf.GetVersions()
			if err != nil{
				return err
			}

			// 定義
			newList := wine.VersionList{}
			errs := []string{}
			

			for _, v := range *vers{
				new, err := wine.GetLocalVersion(v.Cmd) // 現在のバージョンを元に新しく取得
				if err !=nil{
					errs=append(errs, err.Error())
				}

				newList = append(newList, wine.NewVersion(v.Name, v.Cmd, new.No))
			}


			// エラーが1つ以上起きている場合
			if len(errs) >0 {
				return errors.New(strings.Join(errs, "\n"))
			}

			// 書き込み
			err = conf.WriteVersions(&newList)
			return err
		},
	}

	return &cmd
}


