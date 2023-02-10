package main

import (
	"fmt"
	"os"
	wine "github.com/Hayao0819/wine-cellar/go-wine"
)


func ePrintln(str string){
	fmt.Fprintln(os.Stderr, str)
}

func main(){
	homedir, err := os.UserHomeDir()
	winedir := fmt.Sprintf("%s/.wine", homedir)

	if err != nil{
		ePrintln(err.Error())
		os.Exit(1)
	}

	local, err:= wine.GetLocalWine("wine", winedir)
	if err != nil{
		ePrintln(err.Error())
		os.Exit(1)
	}
	fmt.Println(local.Arch.Name)
}
