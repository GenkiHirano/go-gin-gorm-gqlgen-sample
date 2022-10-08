package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf(".envファイルの読み込みが出来ませんでした: %v", err)
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	loadEnv()
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_KEY_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	return api
}

func main() {
	api := getTwitterApi()
	text := "Test Tweet"

	tweet, err := api.PostTweet(text, nil)

	if err != nil {
		panic(err)
	}

	fmt.Print(tweet.Text)
}
