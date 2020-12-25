// main function logic goes here
package pkg

import (
	"log"
	"github.com/dli-invest/finreddit/pkg/login"
	"github.com/jzelinskie/geddit"
	"fmt"
)


func GetPosts() {
	o, err := login.RedditOAuth()
	if err != nil {
		log.Fatal("Failed to initialize Reddit Scrapper")
	}
	log.Println(o)
	if err != nil {
		log.Fatal("Failed to get subreddit submissions")
	}
	subOpts := geddit.ListingOptions{
		Time: "day",
		Limit: 1000,
	}
	// for every subreddit go through and get submissions
	submissions, _ := o.SubredditSubmissions("investing", geddit.NewSubmissions, subOpts)
	for i, s := range submissions {
		fmt.Println(i, s)
		fmt.Println(s.FullID)
	}
	// log.Println(submissions)
}