package api

import "log"

func WebhookToDiploma() error {
	level, skills, err := GetCursusUser("cpieri", "21")
	if err != nil {
		return err
	}
	log.Println(level, skills)
	return nil
}
