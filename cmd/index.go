package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
	"os"
)

func PrintHelp() {
	commands := map[string]string{
		"add":                    "create a new profile",
		"edit":                   "edit an existing profile",
		"su, sp, select, switch": "list all profiles and select one to use",
		"list":                   "list all profiles",
		"delete, del":            "delete a selected profile",
		"current":                "show current active git account",
	}

	fmt.Println("Usage: git-profiles <option>")
	fmt.Println("\nOptions:")
	for cmd, desc := range commands {
		fmt.Printf("  %-30s %s\n", cmd, desc)
	}
}

func App() {
	configFile := config.FindOrCreateConfigFile()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			ExecAddRoutine(configFile)
		case "edit":
			ExecEditRoutine(configFile)
		case "su", "sp", "select", "switch":
			SwitchProfileRoutine(configFile)
		case "list", "ls":
			ListProfiles(configFile)
		case "delete", "del":
			DeleteProfileRoutine(configFile)
		case "current":
			GetCurrentUser()
		case "help", "-h", "--help":
			PrintHelp()
		default:
			PrintHelp()
			fmt.Printf("\n\n\"%s\" is an invalid command.\n", os.Args[1])
		}
	} else {
		PrintHelp()
	}

}
