package main

import (
	"github.com/jzelinskie/geddit"
	"log"
	"github.com/dli-invest/finreddit/pkg/util"
)

func main() {
	client_id := util.GetEnvVar("REDDIT_CLIENT_ID")
	client_secret := util.GetEnvVar("REDDIT_CLIENT_SECRET")
	password := util.GetEnvVar("REDDIT_PASSWORD")
	username := util.GetEnvVar("REDDIT_USERNAME")
	o, err := geddit.NewOAuthSession(
		client_id,
		client_secret,
		"Stonk Market Scrapper see source https://github.com/jzelinskie/geddit",
		"http://friendlyuser.github.io",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create new auth token for confidential clients (personal scripts/apps).
	err = o.LoginAuth(username, password)
	if err != nil {
		log.Fatal(err)
	}

	// Ready to make API calls!
}