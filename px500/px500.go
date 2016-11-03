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
	f := *app.ImageFiles(2)
	page1 := fmt.Sprintf("%v%v", "https://www.500px.com", randPhoto(photos).Url)
	page2 := fmt.Sprintf("%v%v", "https://www.500px.com", randPhoto(photos).Url)
	app.Scrape(page1, f[0])
	app.Scrape(page2, f[1])
	app.ApplyDesktop(f[0], f[1])
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
