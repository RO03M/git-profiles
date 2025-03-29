package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
	"git-profiles/cmd/utils"
	"os"

	"github.com/orochaa/go-clack/prompts"
)

func DeleteProfileRoutine(configFile config.Config) {
	var deleteParam string

	if len(os.Args) >= 3 {
		deleteParam = os.Args[2]
	} else {
		profiles := utils.Map(configFile.Profiles, func(item config.Profile, key int) *prompts.SelectOption[string] {
			return &prompts.SelectOption[string]{
				Label: item.ProfileName,
				Value: item.Email,
				Hint:  fmt.Sprintf("%s %s", item.Name, item.Email),
			}
		})

		deleteParam, _ = prompts.Select(prompts.SelectParams[string]{
			Message: "Choose a profile to be deleted",
			Options: profiles,
		})
	}

	profiles := utils.Filter(configFile.Profiles, func(profile config.Profile, _ int) bool {
		return !(profile.ProfileName == deleteParam || profile.Name == deleteParam || profile.Email == deleteParam)
	})

	configFile.Profiles = profiles
	configFile.Save()
}
