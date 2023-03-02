/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	env "github.com/Hayao0819/wine-cellar/cmd/env"
	version "github.com/Hayao0819/wine-cellar/cmd/version"
)


func init() {
	root.AddCommand(env.EnvCmd)
	root.AddCommand(version.VersionCmd)
}
