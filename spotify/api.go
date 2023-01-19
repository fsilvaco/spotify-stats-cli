package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/fsilvaco/spotify-stats-cli/prompt"
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

type Artists struct {
	Items []struct {
		Name string `json:"name"`
	} `json:"items"`
}

const (
	BASE_URL_API = "https://api.spotify.com/v1/me"
)

var token string

func (s Server) getUserTopItems(token string, search string) {
	var endpoint = BASE_URL_API + "/top/" + search
	client := &http.Client{}

	req, err := http.NewRequest("GET", endpoint, nil)
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

	var artists Artists

	err = json.Unmarshal(body, &artists)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Last 6 months")
	for i := 0; i < len(artists.Items); i++ {
		position := strconv.Itoa(i + 1)
		fmt.Println(position + "- " + artists.Items[i].Name)
	}

}

func (s Server) getCurrentUser(token string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", BASE_URL_API, nil)
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

	result := prompt.Select(user.DisplayName)

	s.getUserTopItems(token, result)

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
