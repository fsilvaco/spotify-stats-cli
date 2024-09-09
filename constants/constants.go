package constants

import "fmt"

const (
	AuthDataDir   = "auth_data"
	TokenFileName = "token.json"
)

func TokenFilePath() string {
	return fmt.Sprintf("%s/%s", AuthDataDir, TokenFileName)
}
