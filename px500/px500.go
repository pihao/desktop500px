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
	p := randPhoto(photos)
	pageUrl := fmt.Sprintf("%v%v", "https://www.500px.com", p.Url)

	app.Scrape(pageUrl)
}

func randPhoto(photos *Photos) *Photo {
	return __randPhoto(photos, 0)
}
func __randPhoto(photos *Photos, randCount int) *Photo {
	p := photos.Photos[rand.Intn(len(photos.Photos))]
	if p.Width > p.Height || randCount > 10 {
		return &p
	} else {
		randCount++
		return __randPhoto(photos, randCount)
	}
}
