package conf

import (
	"path"
)

var (
	ConfDir     string
	EnvConfFile string
	VerConfFile string
	CurrentFile string
)


func Initilize(){
	EnvConfFile = path.Join(ConfDir, "wine-env.json")
	VerConfFile = path.Join(ConfDir, "wine-version.json")
	CurrentFile = path.Join(ConfDir, "current_env")
}
