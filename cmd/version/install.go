package version

import "github.com/spf13/cobra"

func newInstallCmd()*cobra.Command{
	cmd := cobra.Command{
		Use: "install",
		Short: "Wineのインストール",
		Long: `バージョンを指定してWineをローカルにインストールします。`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// https://phoenicis.playonlinux.com/index.php/wine?os=linux
			return nil
		},
	}
	return &cmd
}
