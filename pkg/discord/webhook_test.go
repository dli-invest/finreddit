package discord

import (
	"testing"
	"github.com/dli-invest/finreddit/pkg/types"
)
func TestGetPosts(t *testing.T) {
	discordEmbed := []types.DiscordEmbed{{Title: "sample title", Url: "https://github.com/FriendlyUser/finfiber/blob/master/pkg/discord/discord.go", Description: "CI/CD"}}
	discordPayload := types.DiscordWebhook{Content: "test content", Embeds: discordEmbed}
	_, err := SendWebhook(discordPayload)
	if err != nil {
		t.Errorf("Failed to send request")
	}
}