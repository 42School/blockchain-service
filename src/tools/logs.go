package tools

import (
	"log"
)

func LogsDev(msg string) {
	if Env == "Dev" {
		log.Println("Developper mode:", msg)
	}
}

func LogsError(_err error) {
	log.Println("Error:", _err)
}

func LogsMsg(msg string) {
	log.Println(msg)
}
