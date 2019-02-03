package main

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// https://developer.twitter.com/en/apps
// https://developer.twitter.com/en/docs/tweets/search/api-reference/get-search-tweets

func GetTweets() []twitter.Tweet {
	config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	search, res, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "golang",
	})

	if err != nil {
		panic(err.Error())
	}

	if res.StatusCode >= 300 {
		panic(err.Error())
	}

	return search.Statuses
}
