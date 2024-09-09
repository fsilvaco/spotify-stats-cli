package server

import (
	"fmt"
	"os"

	"github.com/fsilvaco/spotify-stats-cli/token"
	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
)

func Initialize() {
	client_id := os.Getenv("CLIENT_ID")
	port := ":8080"
	scope := "user-top-read user-read-private user-read-email"
	redirect_uri := "http://localhost" + port + "/auth"
	spotify_auth_url := "https://accounts.spotify.com/authorize?" + "client_id=" + client_id + "&response_type=token&scope=" + scope + "&redirect_uri=" + redirect_uri

	r := gin.Default()

	fmt.Printf("Opening the Spotify Login Dialog in your browser...")
	browser.OpenURL(spotify_auth_url)

	r.GET("/auth", func(c *gin.Context) {
		c.File("static/auth.html")

	})

	r.GET("/token", func(c *gin.Context) {
		access_token := c.Query("access_token")
		token_type := c.Query("token_type")
		expires_in := c.Query("expires_in")

		tokenData := token.TokenData{AccessToken: access_token, TokenType: token_type, ExpiresIn: expires_in}

		token.CreateJsonFile(tokenData)

		os.Exit(0)
	})

	r.Run()

}
