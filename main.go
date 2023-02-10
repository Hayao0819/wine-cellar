package main

import (
	"fmt"
	wine "github.com/Hayao0819/wine-cellar/go-wine"
)


func main(){
	local := wine.GetLocalWine
	fmt.Println(local.arch.name)
}
