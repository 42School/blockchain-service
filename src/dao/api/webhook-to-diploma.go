package api

import "log"

func WebhookToDiploma() (diploma, error) {
	level, skills, err := GetCursusUser("cpieri", "21")
	GetBirthdate("cpieri")
	if err != nil {
		return err
	}
	log.Println(level, skills)
	return nil
}
