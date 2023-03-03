package wine

import (
	"fmt"
	"path"
	//"os"
	"os/exec"
	"strconv"
	"strings"
	//"path/filepath"
)


type Version struct{
	Name string
	Cmd string
	No float32
}

type arch struct{
	Name string
}

type Env struct {
	Name string
	Arch arch
	Prefix string
	Version Version
}


func NewVersion(name string, cmd string, no float32) Version {
	return Version{Name: name, Cmd: cmd , No: no}
}

func NewEnv(name string, archStr string, prefix string, version Version) Env {
	arch := *ArchFromStr(archStr)
	return Env{Name: name, Arch: arch, Prefix: prefix, Version: version}
}




func GetLocalVersion(cmd string) (*Version, error){
	var wineVerString string
	var no float64 //strconv.ParseFloatがfloat64のため。あとでfloat32に変換する。
	var err error

	// Get command result
	cmdResult , err := exec.Command(cmd, "--version").Output()
	if err!=nil{
		return nil, ErrFailedExecWine
	}

	// get wine-x.x string
	wineVerString = strings.Split(string(cmdResult), " ")[0]

	// Convert string to float
	//no = float32(strings.Split(name, "-")[1])
	no, err = strconv.ParseFloat(strings.Split(wineVerString, "-")[1], 32)
	if err != nil{
		return nil, ErrFailedConvToFloat
	}


	ver := NewVersion(path.Base(cmd), cmd, float32(no))

	return &ver,nil
}

/*

// ネットの情報だとwinepath -u c:\\windows\\system32の実行結果でsystem32なら32bit、syswow64なら64bitらしかったけどmacOSで動かなかったので無し

func runWinePath(prefix, path string)(string, error){
	runCmd := exec.Command("winepath", "-u", path)
	runCmd.Env = append(os.Environ(), fmt.Sprintf("WINEPREFIX=%s", prefix))
	cmdResult, err := runCmd.Output()
	if err != nil{
		return "",  err //ErrFailedGetPath
	}
	return string(cmdResult), nil
}


func getWineArch(prefix string)(string, error){
	sysdir, err := runWinePath(prefix, "c:\\windows\\system32")
	if err != nil{
		return "", err
	}

	if strings.HasPrefix(sysdir, )
}
*/


// WINEPREFIX/drive_c/windows/syswow64があればwin64、なければwin32
// winetricksも使ってる方法らしいです
func getArch(prefix string)(*arch, error){
	syswow64 := fmt.Sprintf("%s/drive_c/windows/syswow64", prefix)
	system32 := fmt.Sprintf("%s/drive_c/windows/system32", prefix)
	if isDir(syswow64){
		return &Win64, nil
	}else if isDir(system32){
		return &Win32, nil
	}
	return nil, ErrThisDirIsNotPrefix
}


func GetLocalWineEnv (cmd,prefix string)(*Env, error){
	name := "Default wine prefix"
	ver, err := GetLocalVersion(cmd)
	if err!=nil{
		return nil, ErrFailedGetWineVer
	}

	arch, err := getArch(prefix)
	if err != nil{
		return nil, err
	}

	wine := NewEnv(name,arch.Name,prefix,*ver)

	return &wine,nil
}

func ArchFromStr(arch string)(*arch){
	arch = strings.ToLower(arch)
	switch(arch){
		case "win64":
			return &Win64
		case "win32":
			return &Win32
		default:
			return nil
	}
}
