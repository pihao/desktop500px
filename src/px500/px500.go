package px500

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pihao/desktop500px/src/app"
)

func Run() {
	rand.Seed(time.Now().Unix())

	client, accessToken := getClientAndToken(false)
	api := API{*client, *accessToken}

	photos := api.GetPhotos()
	f := *app.GenerateFilePaths(app.GetDesktopCount())
	for i := range f {
		p := fmt.Sprintf("%v%v", "https://www.500px.com", randPhoto(photos).Url)
		app.Scrape(p, f[i])
		app.ApplyDesktop(f[i], i)
	}
	app.Trace("Complete.")
}

func randPhoto(photos *Photos) *Photo {
	return __randPhoto(photos, 0)
}
func __randPhoto(photos *Photos, randTimes int) *Photo {
	p := photos.Photos[rand.Intn(len(photos.Photos))]
	if p.Width > p.Height || randTimes > 10 {
		return &p
	} else {
		randTimes++
		return __randPhoto(photos, randTimes)
	}
}
