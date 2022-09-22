package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	api := getTwitterApi()
	v := url.Values{}
	// トレンドを取得したい地域に日本を指定
	trendResp, err := api.GetTrendsByPlace(23424856, v)

	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()

	fmt.Println(trendResp.AsOf) // 取得実行した時間 (UTC)
	fmt.Println(trendResp.CreatedAt) // トレンドに含まれる一番古い時期からあるものの日時  (UTC)
	fmt.Println(trendResp.Locations) // 指定した地域
	fmt.Println(len(trendResp.Trends)) // トレンドの数
	fmt.Println(t) // 実行時間 (JST)

	// https://pkg.go.dev/github.com/chimeracoder/anaconda#TrendResponse
	// https://pkg.go.dev/github.com/chimeracoder/anaconda#Trend
	for i, v := range trendResp.Trends {
		ranking := i + 1
		fmt.Printf("%d %s\n", ranking, v.Name)
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	// env実装予定
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"),
	os.Getenv("ACCESS_TOKEN_SECRET"))
}
