package conf

import (
	"encoding/json"
	//"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	wine "github.com/Hayao0819/wine-cellar/go-wine"
)

/*
var ErrNoSuchVersion error = errors.New("cant find such version")
var ErrNoSuchEnv error = errors.New("cant find such environment")
*/

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


func readEnv()(*[]env, error){
	wineEnvJson, err := os.Open(EnvConfFile)
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

func readVer()(*[]version, error){
	wineVerJson, err := os.Open(VerConfFile)
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

func GetVersions()(*wine.VersionList, error){
	version , err := readVer()
	if err != nil{
		return nil, err
	}
	var wineVersion wine.VersionList
	for _, ver := range *version{
		wineVersion=append(wineVersion, wine.NewVersion(ver.Name, ver.Cmd, ver.No))
	}

	return &wineVersion, nil
}

/*
func getVerisonFromName(name string, vers *[]wine.Version)(*wine.Version, error){
	for _, ver := range *vers{
		if ver.Name == name{
			return &ver, nil
		}
	}
	return nil, ErrNoSuchVersion
}
*/

func GetEnvs()(*wine.EnvList, error){
	env , err := readEnv()
	if err != nil{
		return nil, err
	}
	

	var wineVer *wine.Version
	var wineEnv wine.Env
	var wineEnvs wine.EnvList

	wineVers , err := GetVersions()
	if err != nil{
		return nil, err
	}

	for _, env := range *env{
		//wineVer, err = getVerisonFromName(env.Version, wineVers)
		wineVer, err =wineVers.GetFromName(env.Version)
		if err != nil{
			fmt.Fprintln(os.Stderr, "Error found in \"" + env.Name + "\"")
			return nil, err

		}
		wineEnv = wine.NewEnv(env.Name, env.Arch, env.Prefix, *wineVer)
		wineEnvs=append(wineEnvs, wineEnv)
	}
	return &wineEnvs,nil
}

func WriteEnvs(wineEnv *wine.EnvList)(error){
	s := []env{}
	for _, w := range *wineEnv{
		e:= env{
			Name: w.Name,
			Arch: w.Arch.Name,
			Prefix: w.Prefix,
			Version: w.Version.Name,
		}
		s = append(s, e)
	}

	json, err := json.Marshal(s)
	if err !=nil{
		return err
	}
	return ioutil.WriteFile(EnvConfFile, json, 0644)

}

func WriteVersions(wineVer *wine.VersionList)error{
	s := []version{}
	for _, w := range *wineVer{
		v:= version{
			Name: w.Name,
			Cmd: w.Cmd,
			No: w.No,
		}
		s = append(s, v)
	}

	json, err := json.Marshal(s)
	if err !=nil{
		return err
	}
	return ioutil.WriteFile(VerConfFile, json, 0644)
}

func SetCurrentEnv(env *wine.Env)(error){
	statusFile, err := os.Create(CurrentFile)
	if err !=nil{
		return err
	}


	if _, err := statusFile.WriteString(env.Name); err !=nil{
		return err
	}

	return nil
}

func GetCurrentEnv()(*wine.Env, error){
	status, err := os.ReadFile(CurrentFile)
	if err !=nil{
		return nil, err
	}

	currentEnvName := strings.Split(string(status), "\n")[0]


	envs, err := GetEnvs()
	if err !=nil{
		return nil, err
	}

	return envs.GetFromName(currentEnvName)
}
