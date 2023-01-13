package main

import (
	"fmt"
	"net/http"

	"github.com/pkg/browser"
)

const (
	PORT         string = ":8080"
	CLIENT_ID    string = "33cfd3aa25f144a18bc39ba4f7b6302c"
	SCOPE        string = "user-top-read"
	REDIRECT_URI string = "http://localhost" + PORT + "/callback"
	AUTH_URL     string = "https://accounts.spotify.com/authorize?" + "client_id=" + CLIENT_ID + "&response_type=token&scope=" + SCOPE + "&redirect_uri=" + REDIRECT_URI
)

func handlerCallback(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "callback.html")
}

func handlerToken(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("access_token")
	fmt.Printf("Your token is: " + token)
}

func main() {
	fmt.Println("Opening the Spotify Login Dialog in your browser...")
	browser.OpenURL(AUTH_URL)
	http.HandleFunc("/callback", handlerCallback)
	http.HandleFunc("/token", handlerToken)
	http.ListenAndServe(PORT, nil)
}
