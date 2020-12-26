package discord 

import (
	"os"
	// "encoding/json"
	"net/http"
	"github.com/dli-invest/finreddit/pkg/types"
	"bytes"
)




func SendWebhook() {
	discordWebhook := new(types.DiscordWebhook)
	discordUrl := os.Getenv("DISCORD_WEBHOOK")
	resp, err := http.Post(discordUrl, "application/json", bytes.NewBuffer(webhookData))
	log.Println(resp)
	log.Println(discordWebhook.Content) // john
	log.Println(discordWebhook.Embeds) // doe
}
