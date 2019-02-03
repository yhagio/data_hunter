package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
)

func showResults(remoteOkJobs []RemoteOKJob, githubJobs []GithubJob, tweets []twitter.Tweet) {
	fmt.Printf("\nRemoteOKJob: %v\n", remoteOkJobs[2])
	fmt.Printf("\nGithubJob: %v\n", githubJobs[2])
	fmt.Printf("\nTweet: %v\n", tweets[2])
}

func main() {
	fmt.Println("---Lets Hunt---")

	out1 := make(chan []RemoteOKJob)
	out2 := make(chan []GithubJob)
	out3 := make(chan []twitter.Tweet)

	go func() {
		out1 <- GetRemoteOKJobs()
	}()

	go func() {
		out2 <- GetGithubJobs()
	}()

	go func() {
		out3 <- GetTweets()
	}()

	showResults(<-out1, <-out2, <-out3)

	os.Exit(0)
}
