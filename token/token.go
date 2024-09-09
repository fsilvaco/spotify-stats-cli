package token

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fsilvaco/spotify-stats-cli/constants"
)

type TokenData struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}

func createDiretory(dir string) {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal("error creating tokens directory")
	}
}

func CreateJsonFile(t TokenData) {

	createDiretory(constants.AuthDataDir)

	file, err := os.Create(constants.TokenFilePath())
	if err != nil {
		log.Fatal("error creating token file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(t); err != nil {
		log.Fatal("error writing token to file")
	}
}

func GetTokenData() TokenData {
	var t TokenData

	file, err := os.Open(constants.TokenFilePath())

	if err != nil {
		log.Fatal("error opening token file")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&t); err != nil {
		log.Fatal("error reading token from file")
	}

	return t
}
