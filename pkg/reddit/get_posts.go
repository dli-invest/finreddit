package reddit 
import (
	"log"
	"github.com/dli-invest/finreddit/pkg/login"
	"github.com/dli-invest/finreddit/pkg/util"
	"github.com/dli-invest/finreddit/pkg/csvs"
	"github.com/dli-invest/finreddit/pkg/types"
	"github.com/jzelinskie/geddit"
	"fmt"
)


// gets submissions a given SRConfiguration
func GetSubmissions(session *geddit.OAuthSession, cfg types.SRConfig) ([]*geddit.Submission) {
	subreddit := cfg.Name
	limit := cfg.Limit
	subOpts := geddit.ListingOptions{
		Limit: limit,
	}
	submissions, err := session.SubredditSubmissions(subreddit, geddit.NewSubmissions, subOpts)
	if err != nil {
		log.Fatal("Failed to retrieve subreddit posts for " + subreddit)
	}
	// further filter entries by minScore and minComments
	var validSubmissions = []*geddit.Submission{}

	for _, submission := range submissions {
        if(submission.NumComments > cfg.MinComments && submission.Score > cfg.MinScore) {
			validSubmissions = append(validSubmissions, submission)
		}
	} 
	return validSubmissions
}

// Scans subreddits from config file
// for example cmd/scan_sr/simple.yml
func ScanSRs(cfgPathStr string) {
	// login to reddit
	o, err := login.RedditOAuth()
	if err != nil {
		log.Fatal("Failed to initialize Reddit Scrapper")
	}
	// read subreddits from config file
	cfgPath := util.MkPathFromStr(cfgPathStr)
	cfg, err := util.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	csvsPath := util.MkPathFromStr("internal/posts.csv")
	for _, srCfg := range cfg.Data.SubReddits {
		srSubmissions := GetSubmissions(o, srCfg)
		for _, s := range srSubmissions {
			// check if submission is in csv already
			// aware that constantly opening the csv is inefficient
			// but I am dealing with a reasonable amount of entires
			hasValue := csvs.FindInCsv(csvsPath, s.FullID, 1)
			if hasValue {
				// if not send to discord
				fmt.Println("subreddit submission already set")
				fmt.Println(s)
			} else {
				// append to csv
				sData := [][]string{{srCfg.Name, s.FullID, s.URL}}
				csvs.AppendToCsv(csvsPath, sData)
				// send to discord
			}
		}
	}
}