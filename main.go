/*
Copyright © 2022 Naz M. <nazlee41@gmail.com>

*/
package main

import (
	"github.com/0xlilnas/cli-app/cmd"
	"github.com/0xlilnas/cli-app/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
