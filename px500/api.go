package px500

import (
	"encoding/json"

	"github.com/pihao/desktop500px/app"
	"github.com/pihao/go-oauth/oauth"
)

type API struct {
	Client      oauth.Client
	AccessToken oauth.Credentials
}

type PX500Error struct {
	Error  string
	Status int
}

func ResponseHasError(body []byte) bool {
	var o PX500Error
	json.Unmarshal(body, &o)
	if (o == PX500Error{}) {
		return false
	} else {
		app.Trace(string(body))
		return true
	}
}
