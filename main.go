package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
)

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

func main() {
	// Json読み込み
	raw, error := ioutil.ReadFile("path/to/twitterAccount.json")

	if error != nil {
		fmt.Println(error.Error())
		return
	}

	var twitterAccount TwitterAccount

	// 構造体にセット
	json.Unmarshal(raw, &twitterAccount)

	// 認証
	// envファイルに置き換える
	anaconda.NewTwitterApiWithCredentials("1254473461310713857-2mg3A9L46BXF5cJma3v7Rvq1sIJFEC", "3iq3muHujMiHCMc1TLxiCrFBlfwS7uS61qdRgeDoKach4", "oIOmfUk3qk4CFNVBMMNQmMeN7", "daaKq64kqwDtiKxS8orUSPqjPhmmZFAc3ogC6GjdVnpkLBajr0")
}
