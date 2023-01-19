package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/browser"
)

type Server struct {
	Port           string
	SpotifyAuthURL string
}

type User struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

var token string

func (s Server) getCurrentUser(token string) {

	url := "https://api.spotify.com/v1/me"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var user User

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Hi, " + user.DisplayName + "\nChoose which information you want to see:")

}

func (s Server) responseAuthStaticFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/auth.html")
}

func (s Server) getTokenUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("access_token")

	if token != "" {
		fmt.Println("Login successful!")
	}

	s.getCurrentUser(token)

}

func (s Server) Server() {
	fmt.Println("Opening the Spotify Login Dialog in your browser...")
	browser.OpenURL(s.SpotifyAuthURL)
	http.HandleFunc("/auth", s.responseAuthStaticFile)
	http.HandleFunc("/token", s.getTokenUser)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}
