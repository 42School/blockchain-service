package tools

import (
	"net/smtp"
)

func getAuth() smtp.Auth {
	auth := smtp.PlainAuth(
		"",
		EmailSender,
		PasswordEmail,
		EmailHost,
	)
	return auth
}

func SendMail(msg string, address string) bool {
	auth := getAuth()
	body := "This Ethereum account " + address + ", no longer has sufficient funds to be able to write other diplomas on the blockchain.\nPlease add more Ethereum to the account.\nAnother account has just taken over writing on the blockchain, if this is not the case the diplomas are put in queue."
	msg = "From:  Blockchain-Service <" + EmailSender + ">\n" + "To: " + ToEmail + "\n" + "Subject: [Blockchain-Service]: Empty Account\n\n" + body
	err := smtp.SendMail("smtp.gmail.com:587", auth, EmailSender, []string{ToEmail}, []byte(msg))
	if err != nil {
		LogsError(err)
		return false
	}
	return true
}
