package api

import (
	"context"
	"encoding/json"
	"github.com/42School/blockchain-service/src/tools"
)

type userCandidature struct {
	BirthDate string `json:"birth_date"`
}

func GetBirthdate(userLogin string) (string, error) {
	var candUser userCandidature
	server := FtApi.Client(context.Background())
	url := "https://api.intra.42.fr/v2/users/" + userLogin + "/user_candidature"
	rest, err := server.Get(url)
	if err != nil {
		tools.LogsError(err)
		return "", err
	}
	err = json.NewDecoder(rest.Body).Decode(&candUser)
	if err != nil {
		tools.LogsError(err)
		return "", err
	}
	return candUser.BirthDate, nil
}
