package constants

import (
	"fmt"
)

const (
	AuthDataDir    = "auth_data"
	TokenFileName  = "token.json"
	SpotifyBaseURL = "https://api.spotify.com/v1/me"
)

func TokenFilePath() string {
	return fmt.Sprintf("%s/%s", AuthDataDir, TokenFileName)
}

func EndpointSpotifyAPI(search string, time string) string {
	return fmt.Sprintf("%s/top/%s?time_range=%s", SpotifyBaseURL, search, time)
}
