package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/PuerkitoBio/goquery"
)

// "https://500px.com/photo/179916915/matera-blue-hour-ii-by-pier-luigi-girola"
func Scrape(pageUrl string) {
	if pageUrl == "" {
		fmt.Println("pageUrl is none.")
		return
	}
	fmt.Println(pageUrl)

	imgUrl := getImageUrl(pageUrl)
	if imgUrl == "" {
		fmt.Println("imgUrl is none.")
		return
	}
	fmt.Println(imgUrl)

	saveImage(imgUrl)
}

// https://drscdn.500px.org/photo/180711743/q%3D80_m%3D2000/90e17dd62445eed988029bdf528906b2
func saveImage(url string) {
	rsp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fileName := formatImageName(url)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer rsp.Body.Close()
	defer file.Close()

	_, err = io.Copy(file, rsp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileName, "\nOK.")
}

func formatImageName(url string) string {
	dir := AppDir + "/"
	name := url[len(url)-32:]
	ext := ".jpg"
	return dir + name + ext
}

func getImageUrl(pageUrl string) string {
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
