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
		case "su", "sp", "select", "switch":
			SwitchProfileRoutine(configFile)
		case "list":
			ListProfiles(configFile)
		case "delete", "del":
			DeleteProfileRoutine(configFile)
		}
	}

}
