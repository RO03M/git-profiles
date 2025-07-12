package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
)

func GetCurrentUser() {
	configFile := config.FindOrCreateConfigFile()

	activeProfileId := configFile.ActiveProfile

	for _, profile := range configFile.Profiles {
		if profile.Id == activeProfileId {
			fmt.Printf("%s: %s, %s\n", profile.ProfileName, profile.Name, profile.Email)
			return
		}
	}

	fmt.Println("No active account. Create one using the command <add>")
}
