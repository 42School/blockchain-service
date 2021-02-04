package api

import (
	"context"
	"encoding/json"
	"github.com/42School/blockchain-service/src/tools"
	"strconv"
)

type skillsApi struct {
	Name	string	`json:"name"`
	Level	float64	`json:"level"`
}

type cursusUser struct {
	Level	float64		`json:"level"`
	Skills	[]skillsApi	`json:"skills"`
}

func GetCursusUser(cursusId int) (float64, []float64, error) {
	var cursus cursusUser
	var skills []float64
	server := FtApi.Client(context.Background())
	url := "https://api.intra.42.fr/v2/cursus_users/" + strconv.Itoa(cursusId)
	rest, err := server.Get(url)
	if err != nil {
		tools.LogsError(err)
		return 0, skills, err
	}
	err = json.NewDecoder(rest.Body).Decode(&cursus)
	if err != nil {
		tools.LogsError(err)
		return 0, skills, err
	}
	level := cursus.Level
	for i := 0; i < len(cursus.Skills); i++ {
		skills = append(skills, cursus.Skills[i].Level)
	}
	return level, skills, nil
}