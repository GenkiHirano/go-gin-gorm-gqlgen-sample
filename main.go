package main

import (
	"github.com/GenkiHirano/go-twitter-api.git/twitter"
)

func main() {
	twitter.PostTweet("テストです")
	twitter.GetSearch("Gopherくん")
	twitter.GetHomeTimeline(50)
}
