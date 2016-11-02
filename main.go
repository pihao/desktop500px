package main

import (
	"fmt"
	"time"
	"math/rand"

	"github.com/pihao/desktop500px/app"
	"github.com/pihao/desktop500px/px500"
)

const (
	version = `desktop500px 0.1`
)

func main() {
	rand.Seed(time.Now().Unix())

	client, accessToken := px500.GetClientAndToken(false)
	api := px500.API{*client, *accessToken}

	photos := api.GetPhotos()
	i := rand.Intn(len(photos.Photos))
	pageUrl := fmt.Sprintf("%v%v", "https://www.500px.com", photos.Photos[i].Url)

	app.Scrape(pageUrl)
}
