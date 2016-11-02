package px500

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pihao/desktop500px/app"
)

func Run() {
	rand.Seed(time.Now().Unix())

	client, accessToken := getClientAndToken(false)
	api := API{*client, *accessToken}

	photos := api.GetPhotos()
	i := rand.Intn(len(photos.Photos))
	pageUrl := fmt.Sprintf("%v%v", "https://www.500px.com", photos.Photos[i].Url)

	app.Scrape(pageUrl)
}
