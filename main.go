package main

import (
	"github.com/common-creation/fld/commands"
)

func main() {
	rootCmd := commands.NewRootCmd()
	_ = rootCmd.Execute()
}
