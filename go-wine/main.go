package wine

import (
	"fmt"
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

type Wine struct {
	Name string
	Arch arch
	Prefix string
	Version Version
}


func NewVersion(name string, cmd string, no float32) Version {
	return Version{Name: name, Cmd: cmd , No: no}
}

func NewWine(name string, arch arch, prefix string, version Version) Wine {
	return Wine{Name: name, Arch: arch, Prefix: prefix, Version: version}
}




func GetLocalVersion(cmd string) (*Version, error){
	var name string
	var no float64 //strconv.ParseFloatがfloat64のため。あとでfloat32に変換する。
	var err error

	// Get command result
	cmdResult , err := exec.Command("wine", "--version").Output()
	if err!=nil{
		return nil, ErrFailedExecWine
	}

	// get wine-x.x string
	name = strings.Split(string(cmdResult), " ")[0]

	// Convert string to float
	//no = float32(strings.Split(name, "-")[1])
	no, err = strconv.ParseFloat(strings.Split(name, "-")[1], 32)
	if err != nil{
		return nil, ErrFailedConvToFloat
	}


	ver := NewVersion(name, cmd, float32(no))

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
func getWineArch(prefix string)(*arch, error){
	syswow64 := fmt.Sprintf("%s/drive_c/windows/syswow64", prefix)
	system32 := fmt.Sprintf("%s/drive_c/windows/system32", prefix)
	if isDir(syswow64){
		return &Win64, nil
	}else if isDir(system32){
		return &Win32, nil
	}
	return nil, ErrThisDirIsNotPrefix
}


func GetLocalWine (cmd,prefix string)(*Wine, error){
	name := "Default wine prefix"
	ver, err := GetLocalVersion(cmd)
	if err!=nil{
		return nil, ErrFailedGetWineVer
	}

	arch, err := getWineArch(prefix)
	if err != nil{
		return nil, err
	}

	wine := NewWine(name,*arch,prefix,*ver)

	return &wine,nil
}
