package main

import (
	"github.com/loksonarius/gli/cmd"
)

// I really wish this variable could've just been set under the cmd pkg >:V
var version = "unset"

func main() {
	cmd.Execute(version)
}
