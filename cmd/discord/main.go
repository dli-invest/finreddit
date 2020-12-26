package main

import (
	"github.com/dli-invest/finreddit/pkg/discord"
	"github.com/dli-invest/finreddit/pkg/types"
	"log"
)

func main() {
	discordEmbed := []types.DiscordEmbed{{Title: "sample title", Url: "https://github.com/FriendlyUser/finfiber/blob/master/pkg/discord/discord.go", Description: "CI/CD"}}
	discordPayload := types.DiscordWebhook{Content: "test content", Embeds: discordEmbed}
	_, err := discord.SendWebhook(discordPayload)
	if err != nil {
		log.Println("Failed to send request")
	}
}