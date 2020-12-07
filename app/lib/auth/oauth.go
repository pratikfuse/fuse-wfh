package auth

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Credentials struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IdToken      string `json:"id_token"`
}

func GetAccessToken(code string, wg *sync.WaitGroup) []byte {
	oauthUrl := "https://oauth2.googleapis.com/token"
	postBody := fmt.Sprintf(`{"client_id": "%s", "client_secret": "%s", code: "%s", "grant_type": "%s", "redirect_uri": "%s"}`,
		"CLIENT_ID",
		"CLIENT_SECRET",
		code,
		"authorization_code",
		"http://localhost:4000/oauth",
	)

	var jsonStr = []byte(postBody)

	request, err := http.NewRequest("POST", oauthUrl, bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Fatal("error getting access token", err)
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal("error retrieving access code ", err)
	}

	respBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal("error parsing oauth response", err)
	}

	wg.Done()
	return respBody
}
