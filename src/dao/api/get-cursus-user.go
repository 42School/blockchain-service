package api

import (
	"context"
	"encoding/json"
	"github.com/42School/blockchain-service/src/tools"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Skill struct {
	Name  string  `json:"name"`
	Level float64 `json:"level"`
}

type cursusUser struct {
	Level  float64          `json:"level"`
	Skills []Skill `json:"skills"`
}

func GetCursusUser(cursusId int) (float64, []Skill, error) {
	var cursus cursusUser
	server := FtApi.Client(context.Background())
	url := "https://api.intra.42.fr/v2/cursus_users/" + strconv.Itoa(cursusId)
	rest, err := server.Get(url)
	if err != nil {
		tools.LogsError(err)
		return 0, []Skill{}, err
	}
	err = json.NewDecoder(rest.Body).Decode(&cursus)
	if err != nil {
		tools.LogsError(err)
		return 0, []Skill{}, err
	}
	level := cursus.Level
	log.Info(cursus.Skills)
	return level, cursus.Skills, nil
}
