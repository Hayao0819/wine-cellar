package conf

import (
	"path"
)

var (
	ConfDir     string
	EnvConfFile string
	VerConfFile string
)


func Initilize(){
	EnvConfFile = path.Join(ConfDir, "wine-env.json")
	VerConfFile = path.Join(ConfDir, "wine-version.json")
}
