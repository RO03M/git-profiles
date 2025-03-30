package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
	"git-profiles/cmd/utils"

	"github.com/orochaa/go-clack/prompts"
)

func ExecEditRoutine(configFile config.Config) {
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

	targetProfile, _ := prompts.Select(prompts.SelectParams[config.Profile]{
		Message: "Select a profile to edit",
		Options: profiles,
	})

	name, _ := prompts.Text(prompts.TextParams{
		Message:      "Git account name",
		Placeholder:  "Mona Lisa",
		InitialValue: targetProfile.Name,
	})

	email, _ := prompts.Text(prompts.TextParams{
		Message:      "Git account email",
		Placeholder:  "yourgitaccount@email.com",
		InitialValue: targetProfile.Email,
	})

	absoluteSshPath, _ := prompts.Path(prompts.PathParams{
		Message:      "Write the path to the ssh credential",
		Required:     true,
		InitialValue: targetProfile.AbsoluteSshPath,
	})

	profileName, _ := prompts.Text(prompts.TextParams{
		Message:      "What should this git profile be called?",
		Required:     false,
		Placeholder:  fmt.Sprintf("Profile for %s", email),
		InitialValue: targetProfile.ProfileName,
	})

	targetProfile.Name = name
	targetProfile.Email = email
	targetProfile.AbsoluteSshPath = absoluteSshPath
	targetProfile.ProfileName = profileName

	configFile.Profiles = utils.Map(configFile.Profiles, func(profile config.Profile, key int) config.Profile {
		if profile.Id != targetProfile.Id {
			return profile
		}

		return targetProfile
	})

	configFile.Save()
}
