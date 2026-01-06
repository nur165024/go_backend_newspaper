package main

import (
	"gin-quickstart/cmd"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		cmd.Migrate()
		return	
	}

	cmd.Server()
}
