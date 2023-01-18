package spotify

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/browser"
)

type Server struct {
	Port           string
	SpotifyAuthURL string
}

var token string

func (s Server) responseAuthStaticFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/auth.html")
}

func (s Server) getTokenUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("access_token")
	fmt.Printf(token)
}

func (s Server) Server() {
	fmt.Println("Opening the Spotify Login Dialog in your browser...")
	browser.OpenURL(s.SpotifyAuthURL)
	http.HandleFunc("/auth", s.responseAuthStaticFile)
	http.HandleFunc("/token", s.getTokenUser)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}
