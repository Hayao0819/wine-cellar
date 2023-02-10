package conf

import (
	"encoding/json"
	"errors"
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

func GetVersions(jsonPath string)(*[]wine.Version, error){
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

func getVerisonFromName(name string, vers *[]wine.Version)(*wine.Version, error){
	for _, ver := range *vers{
		if ver.Name == name{
			return &ver, nil
		}
	}
	return nil, errors.New("cant find such name")
}

func GetEnvs(envJson, verJson string)(*[]wine.Env, error){
	env , err := readEnv(envJson)
	if err != nil{
		return nil, err
	}
	

	var wineVer *wine.Version
	var wineEnv wine.Env
	var wineEnvs []wine.Env

	wineVers , err := GetVersions(verJson)
	if err != nil{
		return nil, err
	}

	for _, env := range *env{
		wineVer, err = getVerisonFromName(env.Name, wineVers)
		if err != nil{
			return nil, err
		}
		wineEnv = wine.NewEnv(env.Name, env.Arch, env.Prefix, *wineVer)
		wineEnvs=append(wineEnvs, wineEnv)
	}
	return &wineEnvs,nil
}
