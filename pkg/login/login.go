package login 

import (
	"github.com/jzelinskie/geddit"
	"log"
	"github.com/dli-invest/finreddit/pkg/util"
)

// returns oauth session for reddit
// or fails - fine for my purposes
func RedditOAuth() (*geddit.OAuthSession, error) {
	client_id := util.GetEnvVar("REDDIT_CLIENT_ID")
	client_secret := util.GetEnvVar("REDDIT_CLIENT_SECRET")
	password := util.GetEnvVar("REDDIT_PASSWORD")
	username := util.GetEnvVar("REDDIT_USERNAME")
	o, err := geddit.NewOAuthSession(
		client_id,
		client_secret,
		"Stonk Market Scrapper see source https://github.com/dli-invest/finreddit",
		"http://friendlyuser.github.io",
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create new auth token for confidential clients (personal scripts/apps).
	err = o.LoginAuth(username, password)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return o, nil
}