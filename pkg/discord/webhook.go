package discord 

import (
	"os"
	"log"
	"encoding/json"
	"net/http"
	"github.com/dli-invest/finreddit/pkg/types"
	"bytes"
)

func SendWebhook(discordWebhook types.DiscordPayload) (*http.Response, error){
	discordUrl := os.Getenv("DISCORD_WEBHOOK")
	if discordUrl == "" {
		log.Fatal("DISCORD_WEBHOOK not set")
	}
	webhookData, err := json.Marshal(discordWebhook)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(discordUrl, "application/json", bytes.NewBuffer(webhookData))
	return resp, err
}
