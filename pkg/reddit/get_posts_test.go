package reddit

import (
	"testing"
	"github.com/dli-invest/finreddit/pkg/login"
	"github.com/dli-invest/finreddit/pkg/util"
)

func TestGetPosts(t *testing.T) {
	o, err := login.RedditOAuth()
	if err != nil {
		t.Errorf("Failed to initialize Reddit Scrapper")
	}
	// read subreddits from config file
	cfgPath := util.MkPathFromStr("sample_cfg.yml")
	cfg, err := util.NewConfig(cfgPath)
	if err != nil {
		t.Errorf("Failed to initialize configuration")
	}
	subreddit := cfg.Data.SubReddits[0]
	submissions := GetSubmissions(o, subreddit)
	
	if len(submissions) < 0 {
		t.Errorf("Expected non zero results")
	}
}
