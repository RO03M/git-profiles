package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
	"git-profiles/cmd/utils"
	"log"

	"github.com/orochaa/go-clack/prompts"
)

func SwitchProfileRoutine(configFile config.Config) {
	profiles := utils.Map(configFile.Profiles, func(item config.Profile, key int) *prompts.SelectOption[config.Profile] {
		return &prompts.SelectOption[config.Profile]{
			Label: item.ProfileName,
			Value: item,
			Hint:  fmt.Sprintf("%s %s", item.Name, item.Email),
		}
	})

	if len(profiles) == 0 {
		prompts.Info("No profiles found, create one using the add command")
		return
	}

	targetProfile, err := prompts.Select(prompts.SelectParams[config.Profile]{
		Options: profiles,
	})

	if err != nil {
		log.Fatalf("Failed to select a profile %v\n", err)
		return
	}

	fmt.Println(targetProfile)
}
