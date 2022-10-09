package twitter

import (
	"fmt"
	"net/url"
	"strconv"
)

// ツイート機能
func PostTweet(text string) {
	funcStart("ツイート機能")
	api := getTwitterApi()

	tweet, err := api.PostTweet(text, nil)

	if err != nil {
		fmt.Printf("failed to tweet: %v\n", err)
	}

	fmt.Println(tweet.Text)

	funcFinish()
}

// 検索機能
func GetSearch(word string) {
	funcStart("検索機能")
	api := getTwitterApi()

	searchResult, err := api.GetSearch(word, nil)

	if err != nil {
		fmt.Printf("failed to get search: %v\n", err)
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}

	funcFinish()
}

// ホームタイムライン取得機能 (最新ツイートから取得)
func GetHomeTimeline(quantity int) {
	funcStart("ホームタイムライン取得機能")
	api := getTwitterApi()

	v := url.Values{}
	s := strconv.Itoa(quantity)

	v.Set("count", s)
	tweets, err := api.GetHomeTimeline(v)

	if err != nil {
		fmt.Printf("failed to get home timeline: %v\n", err)
	}

	for i, tweet := range tweets {
		fmt.Printf("%d, tweetname:%v\n tweet: %v\n", i, tweet.User.Name, tweet.FullText)
	}

	funcFinish()
}

func funcStart(funcName string) {
	fmt.Println(funcName)
	fmt.Println("==================================================")
}

func funcFinish() {
	fmt.Println("==================================================")
	fmt.Println()
}
