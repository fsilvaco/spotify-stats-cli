package login

import (
	"fmt"

	"github.com/fsilvaco/spotify-stats-cli/env"
)

func Inicialize() {
	env.Load()
	port := env.Get("PORT", "8080")
	clientId := env.Get("CLIENT_ID", "")
	fmt.Println("Here the port is", port, clientId)

}
