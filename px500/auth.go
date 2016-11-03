package px500

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"

	"github.com/pihao/desktop500px/app"
	"github.com/pihao/go-oauth/oauth"
)

var accessTokenFile = path.Join(app.AppDir, "access_token.json")

type Key struct {
	ConsumerKey    string
	ConsumerSecret string
	Username       string
	Password       string
}

func getClientAndToken(forceNewToken bool) (*oauth.Client, *oauth.Credentials) {
	client := oauth.Client{
		TemporaryCredentialRequestURI: "https://api.500px.com/v1/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.500px.com/v1/oauth/authorize",
		TokenRequestURI:               "https://api.500px.com/v1/oauth/access_token",
	}
	key := readKey()
	if key == nil {
		return nil, nil
	}
	client.Credentials.Token = key.ConsumerKey
	client.Credentials.Secret = key.ConsumerSecret
	accessToken := getToken(&client, *key, forceNewToken)
	return &client, accessToken
}

func getToken(client *oauth.Client, key Key, forceNewToken bool) *oauth.Credentials {
	accessToken := readAccessToken()
	if forceNewToken || accessToken == nil {
		accessToken = generateAccessToken(client, key)
	}
	return accessToken
}

func generateAccessToken(client *oauth.Client, key Key) *oauth.Credentials {
	requestToken, err := client.RequestTemporaryCredentials(nil, "oob", nil)
	if err != nil {
		app.Trace("ERROR: requestToken error. Check your key file:", app.KeyFile)
		log.Fatal(err)
	}
	accessToken, _, err := client.RequestTokenXAuth(nil, requestToken, key.Username, key.Password)
	if err != nil {
		app.Trace("ERROR: accessToken error.")
		log.Fatal(err)
	}
	saveAccessToken(accessToken)
	return accessToken
}

func saveAccessToken(accessToken *oauth.Credentials) {
	b, err := json.Marshal(*accessToken)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(accessTokenFile, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	app.Trace("Generated new access token.")
}

func readAccessToken() *oauth.Credentials {
	b, err := ioutil.ReadFile(accessTokenFile)
	if err != nil {
		return nil
	}
	var o oauth.Credentials
	json.Unmarshal(b, &o)
	if (o == oauth.Credentials{}) {
		app.Trace("Format Error: %v\n", accessTokenFile)
		return nil
	}
	return &o
}

func readKey() *Key {
	b, err := ioutil.ReadFile(app.KeyFile)
	if err != nil {
		log.Fatal(err)
	}
	var o Key
	json.Unmarshal(b, &o)
	return &o
}
