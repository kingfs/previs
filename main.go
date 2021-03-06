package main

import (
	"fmt"
	"os"

	"github.com/PaulRosset/previs/api"
)

func whichConfig(args []string) string {
	if len(args) == 1 && args[0] == "-p" {
		fmt.Println("You are using '.previs.yml' file")
		return ".previs.yml"
	}
	fmt.Println("You are using '.travis.yml' file")
	return ".travis.yml"
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encountered: %+v\n", err)
		os.Exit(2)
	}
	configFile := whichConfig(os.Args[1:])
	imgDocker, envsVar, err := api.Writter(cwd + "/" + configFile)
	if err != nil {
		api.CleanUnusedDockerfile(cwd, imgDocker)
		fmt.Fprintf(os.Stderr, "error encountered: %+v\n", err)
		os.Exit(2)
	}
	err = api.Start(imgDocker, cwd+"/", envsVar)
	if err != nil {
		api.CleanUnusedDockerfile(cwd, imgDocker)
		fmt.Fprintf(os.Stderr, "error encountered: %+v\n", err)
		os.Exit(2)
	}
}
