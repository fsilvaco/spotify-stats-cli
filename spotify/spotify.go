package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fsilvaco/spotify-stats-cli/constants"
	"github.com/fsilvaco/spotify-stats-cli/token"
)

type SpotifyResponse struct {
	Items []struct {
		Name string `json:"name"`
	} `json:"items"`
}

var data SpotifyResponse

func GetTopItems(item string) SpotifyResponse {
	c := &http.Client{}
	e := constants.EndpointSpotifyAPI(item, "long_term")
	t := token.GetTokenData()

	authHeader := fmt.Sprintf("%s %s", t.TokenType, t.AccessToken)

	req, err := http.NewRequest("GET", e, nil)
	if err != nil {
		log.Fatal("Erro with request spotify API")
	}

	req.Header.Add("Authorization", authHeader)

	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("Failed to execute the request: %v\n", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error reading response body: %v\n", err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON (body: %s): %v\n", string(body), err)
	}

	return data

}
