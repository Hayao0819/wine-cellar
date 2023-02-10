package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"

	wine "github.com/Hayao0819/wine-cellar/go-wine"
)

type version struct{
	Name string `json:"name"`
	Cmd string `json:"cmd"`
	No float32 `json:"version"`
}

type env struct {
	Name string `json:"name"`
	Arch string `json:"arch"`
	Prefix string `json:"prefix"`
	Version string `json:"version"`
}


func readEnv(jsonPath string)(*[]env, error){
	wineEnvJson, err := os.Open(jsonPath)
	if err != nil{
		return nil, err
	}

	defer wineEnvJson.Close()

	wineEnvBytes, err := ioutil.ReadAll(wineEnvJson)
	if err != nil{
		return nil, err
	}

	var wineEnv []env
	json.Unmarshal(wineEnvBytes, &wineEnv)

	return &wineEnv, nil
}

func readVer(jsonPath string)(*[]version, error){
	wineVerJson, err := os.Open(jsonPath)
	if err != nil{
		return nil, err
	}

	defer wineVerJson.Close()

	wineEnvBytes, err := ioutil.ReadAll(wineVerJson)
	if err != nil{
		return nil, err
	}

	var wineVer []version
	json.Unmarshal(wineEnvBytes, &wineVer)

	return &wineVer, nil
}

func GetWineVersions(jsonPath string)(*[]wine.Version, error){
	version , err := readVer(jsonPath)
	if err != nil{
		return nil, err
	}
	var wineVersion []wine.Version
	for _, ver := range *version{
		wineVersion=append(wineVersion, wine.NewVersion(ver.Name, ver.Cmd, ver.No))
	}

	return &wineVersion, nil
}

func GetWineEnv(jsonPath string)(*[]wine.Env, error){
	// Todo
	return nil,nil
}
