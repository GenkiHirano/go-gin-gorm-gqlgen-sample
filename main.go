package main

import (
	"github.com/GenkiHirano/go-twitter-api.git/twitter"
)

func main() {
	twitter.PostTweet("テストツイートです")
	twitter.GetSearch("テストツイート")
	twitter.GetHomeTimeline(50)
}
