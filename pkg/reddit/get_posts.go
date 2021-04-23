package reddit 
import (
	"log"
	"github.com/dli-invest/finreddit/pkg/login"
	"github.com/dli-invest/finreddit/pkg/util"
	"github.com/dli-invest/finreddit/pkg/csvs"
	"github.com/dli-invest/finreddit/pkg/types"
	"github.com/dli-invest/finreddit/pkg/discord"
	"github.com/jzelinskie/geddit"
	"strings"
	"strconv"
	"fmt"
	"time"
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
		if (submission.NumComments != 0 && cfg.MinScore != 0) {
			if(submission.NumComments >= cfg.MinComments && submission.Score >= cfg.MinScore) {
				validSubmissions = append(validSubmissions, submission)
				continue
			}
		}
		if (cfg.LinkFlairText != "") {
			// checking for flair
			if (strings.Contains(submission.LinkFlairText, cfg.LinkFlairText)) {
				validSubmissions = append(validSubmissions, submission)
				continue
			}
		}
		if (len(cfg.Phrases) != 0) {
			// search through phrases
			title := strings.ToLower(submission.Title)
				// check matches word
			for _, phrase := range cfg.Phrases {
				// check if phrase is contained in title
				lowerPhrase := strings.ToLower(phrase)
				addSubmission := strings.Contains(title, lowerPhrase)
				if addSubmission {
					validSubmissions = append(validSubmissions, submission)
					continue
				}
			}
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
	csvsPath := util.MkPathFromStr(cfg.Data.CsvPath)
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
				// seems like a lot of posts, wondering if I will hit 
				// post limit, sleep 2 seconds after each post.
				// append to csv
				sData := [][]string{{srCfg.Name, s.FullID, s.URL}}
				csvs.AppendToCsv(csvsPath, sData)
				discordPayload := MapSubmissionToEmbed(s)
				_, err := discord.SendWebhook(discordPayload)
				if err != nil {
					fmt.Println(s.FullID)
					fmt.Println(err)
				}
				time.Sleep(2 * time.Second)
			}
		}
	}
}


func MapSubmissionToEmbed(submission *geddit.Submission)  (types.DiscordPayload) {
	description := fmt.Sprintf(
		"%s (%d Likes, %d Comments)",
		submission.Author,
		submission.Score,
		submission.NumComments)
	// get timestamp 
	var dateCreated int64 = int64(submission.DateCreated)
	t := time.Unix(dateCreated, 0)
	timestamp := t.Format(time.RFC3339)
	title := fmt.Sprintf("%s - %s", submission.Subreddit, submission.Title)
	discordEmbed := []types.DiscordEmbed{{
		Title: title,
		Url: submission.URL,
		Description: description,
		Timestamp: timestamp,
	}}
	discordPayload := types.DiscordPayload{Embeds: discordEmbed}
	return discordPayload
}