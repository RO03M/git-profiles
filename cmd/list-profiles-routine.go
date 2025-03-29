package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
)

func ListProfiles(configFile config.Config) {
	for _, profile := range configFile.Profiles {
		fmt.Printf("%s\t %s\t %s\t %s\t\n", profile.ProfileName, profile.Name, profile.Email, profile.AbsoluteSshPath)
	}
}
