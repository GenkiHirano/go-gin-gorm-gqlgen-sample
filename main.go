package main

import (
	"github.com/ChimeraCoder/anaconda"
)

func main() {
	// twitter api の情報に書き換える
	anaconda.NewTwitterApiWithCredentials("your-access-token", "your-access-token-secret", "your-consumer-key", "your-consumer-secret")
}
