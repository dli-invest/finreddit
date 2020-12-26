package main

import (
	"github.com/dli-invest/finreddit/pkg/discord"
	"github.com/dli-invest/finreddit/pkg/types"
)

func main() {
	discordEmbed := []types.DiscordEmbed{{Title: "sample title", Url: "https://github.com/FriendlyUser/finfiber/blob/master/pkg/discord/discord.go", Description: "CI/CD"}}
	discordPayload := types.DiscordWebhook{Content: "test content", Embeds: discordEmbed}
	discord.SendWebhook(discordPayload)
}