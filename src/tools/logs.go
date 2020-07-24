package tools

import (
	"github.com/42School/blockchain-service/src/global"
	"log"
)

func LogsDev(msg string) {
	if global.Env == "Dev" {
		log.Println("Developper mode:", msg)
	}
}

func LogsError(_err error) {
	log.Println("Error:", _err)
}

func LogsMsg(msg string) {
	log.Println(msg)
}
