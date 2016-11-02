package px500

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Photos struct {
	CurrentPage int
	TotalPages  int
	TotalItems  int
	Photos      []Photo
}

type Photo struct {
	ID          int
	UserID      int `json:"user_id"`
	Name        string
	Camera      string
	Lens        string
	FocalLength string `json:"focal_length"`
	ISO         string
	TimesViewed int `json:"times_viewed"`
	Rating      float32
	Url         string
}

func (api *API) GetPhotos() *Photos {
	rsp, err := api.Client.Get(nil, &api.AccessToken, "https://api.500px.com/v1/photos.json", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if ResponseHasError(body) {
		return nil
	} else {
		var o Photos
		json.Unmarshal(body, &o)
		return &o
	}
}
