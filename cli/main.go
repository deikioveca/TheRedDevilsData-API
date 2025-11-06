package main

import "github.com/deikioveca/TheRedDevilsData/cli/app"

func main() {
	cli 	:= app.NewCLI()
	root 	:= cli.RootCmd()
	root.Execute()
}