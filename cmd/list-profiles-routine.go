package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
)

func ListProfiles(configFile config.Config) {
	for _, profile := range configFile.Profiles {
		fmt.Printf("%s\n  %s\n  %s\n  %s\n \n", profile.ProfileName, profile.Name, profile.Email, profile.AbsoluteSshPath)
	}
}
