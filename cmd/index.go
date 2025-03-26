package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
)

func App() {
	configFile := config.FindOrCreateConfigFile()

	fmt.Println(configFile)
}
