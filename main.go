package main

import (
	"github.com/fsilvaco/spotify-stats-cli/spotify"
)

const (
	PORT             string = ":8080"
	CLIENT_ID        string = "33cfd3aa25f144a18bc39ba4f7b6302c"
	SCOPE            string = "user-top-read"
	REDIRECT_URI     string = "http://localhost" + PORT + "/auth"
	SPOTIFY_AUTH_URL string = "https://accounts.spotify.com/authorize?" + "client_id=" + CLIENT_ID + "&response_type=token&scope=" + SCOPE + "&redirect_uri=" + REDIRECT_URI
)

func main() {
	server := spotify.Server{Port: PORT, SpotifyAuthURL: SPOTIFY_AUTH_URL}
	server.Server()
}
