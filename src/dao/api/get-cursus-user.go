package api

import (
	"context"
	"encoding/json"
)

type skillsApi struct {
	Name string `json:"name"`
	Level float64 `json:"level"`
}

type cursusUser struct {
	Level float64 `json:"level"`
	Skills []skillsApi `json:"skills"`
}

func GetCursusUser(userLogin string, cursusId string) (float64, []float64, error) {
	var cursus []cursusUser
	var skills []float64
	server := FtApi.Client(context.Background())
	url := "https://api.intra.42.fr/v2/users/" + userLogin + "/cursus_users?cursus_id=" + cursusId
	rest, err := server.Get(url)
	if err != nil {
		return 0, skills, err
	}
	err = json.NewDecoder(rest.Body).Decode(&cursus)
	if err != nil {
		return 0, skills, err
	}
	level := cursus[0].Level
	for i := 0; i < len(cursus[0].Skills); i++ {
		skills = append(skills, cursus[0].Skills[i].Level)
	}
	return level, skills, nil
}