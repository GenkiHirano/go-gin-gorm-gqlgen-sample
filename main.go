package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf(".envファイルの読み込みが出来ませんでした: %v", err)
	}

	twitterAccount := &TwitterAccount{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	// 認証
	api := anaconda.NewTwitterApiWithCredentials(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret, twitterAccount.ConsumerKey, twitterAccount.ConsumerSecret)

	// 検索
	searchResult, err := api.GetSearch(`テスト`, nil)

	if err != nil {
		fmt.Println(err)
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}
}
