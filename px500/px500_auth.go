package px500

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/pihao/go-oauth/oauth"
	"github.com/pihao/desktop500px/app"
)

var keyFile = app.AppDir + "/key.json"
var accessTokenFile = app.AppDir + "/access_token.json"

type Key struct {
	ConsumerKey    string
	ConsumerSecret string
	Username       string
	Password       string
}

func GetClientAndToken(forceNewToken bool) (*oauth.Client, *oauth.Credentials) {
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
		log.Fatal(err)
	}
	accessToken, _, err := client.RequestTokenXAuth(nil, requestToken, key.Username, key.Password)
	if err != nil {
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
	fmt.Println("Generated new access token.")
}

func readAccessToken() *oauth.Credentials {
	b, err := ioutil.ReadFile(accessTokenFile)
	if err != nil {
		fmt.Printf("no such file or directory: %v\n", accessTokenFile)
		return nil
	}
	var o oauth.Credentials
	json.Unmarshal(b, &o)
	if (o == oauth.Credentials{}) {
		fmt.Printf("Format Error: %v\n", accessTokenFile)
		return nil
	}
	return &o
}

func readKey() *Key {
	b, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatal(err)
	}
	var o Key
	json.Unmarshal(b, &o)
	return &o
}
