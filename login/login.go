package login

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fsilvaco/spotify-stats-cli/env"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var (
	spotifyOAuthConfig = &oauth2.Config{
		ClientID:     env.Get("SPOTIFY_CLIENT_ID"),
		ClientSecret: env.Get("SPOTIFY_SECRET_ID"),
		RedirectURL:  "http://localhost:8080/auth",
		Scopes:       []string{"user-top-read", "user-read-private", "user-read-email"},
		Endpoint:     spotify.Endpoint,
	}
)

func Inicialize() {
	url := spotifyOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	browser.OpenURL(url)
	http.HandleFunc("/auth", handleAuth)

	// Iniciando o servidor HTTP em uma goroutine
	go func() {
		fmt.Println("Iniciando servidor de autenticação...")
		http.ListenAndServe(":8080", nil)
	}()

	// Definindo um temporizador para encerrar o servidor após um determinado período
	timer := time.NewTimer(3 * time.Second) // O servidor após 3 segundos
	<-timer.C

	fmt.Println("Encerrando o servidor...")
	os.Exit(0)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	token, err := spotifyOAuthConfig.Exchange(r.Context(), code)

	if err != nil {
		fmt.Printf("Erro ao trocar o código de autorização por access_token")
		return
	}
	html := `<html><body><h2>Login efetuado com sucesso!</h2><p>Você pode fechar esta janela</p></body></html>`
	fmt.Fprint(w, html)
	fmt.Printf("Token de acesso: %s\n", token.AccessToken)

}
