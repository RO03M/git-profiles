package cmd

import (
	"git-profiles/cmd/config"
	"os"
)

func App() {
	configFile := config.FindOrCreateConfigFile()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			ExecAddRoutine(configFile)
		case "select":
		case "switch":
		case "sp":
		case "su":
			SwitchProfileRoutine(configFile)
		case "list":
			ListProfiles(configFile)
		}
	}

}
