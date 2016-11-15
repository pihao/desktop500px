package app

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// "https://500px.com/photo/179916915/matera-blue-hour-ii-by-pier-luigi-girola"
func Scrape(pageUrl string, picFile string) {
	if pageUrl == "" {
		log.Fatal("pageUrl is none.")
	}
	Trace("Page:", pageUrl)

	picUrl := getPictureUrl(pageUrl)
	if picUrl == "" {
		log.Fatal("picUrl is none.")
	}
	Trace("Picture:", picUrl)

	savePicture(picUrl, picFile)
}

// https://drscdn.500px.org/photo/180711743/q%3D80_m%3D2000/90e17dd62445eed988029bdf528906b2
func savePicture(url string, file string) {
	rsp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}

	defer rsp.Body.Close()
	defer f.Close()

	_, err = io.Copy(f, rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	Trace("File:", file)
}

func getPictureUrl(pageUrl string) string {
	doc, err := goquery.NewDocument(pageUrl)
	if err != nil {
		log.Fatal(err)
	}
	metaList := doc.Find("head meta")
	for i := range metaList.Nodes {
		e := metaList.Eq(i)
		p, _ := e.Attr("property")
		if p == "og:image" {
			url, _ := e.Attr("content")
			if url != "" {
				return url
			}
		}
	}
	return ""
}
