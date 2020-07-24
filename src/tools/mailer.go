package tools

import "net/smtp"

var emailSender string = "emailSender"
var passwordEmail string = "password"

func getAuth() smtp.Auth {
	auth := smtp.PlainAuth(
		"",
		emailSender,
		passwordEmail,
		"smtp.gmail.com",
	)
	return auth
}

func SendMail(msg string, to string, address string) bool {
	auth := getAuth()
	if msg == "Security Alert" {
		body := "A different hash than the request has just been written in the smart-contract. \nThis is a serious error that compromises the security of the service. \nA security mode has just been activated, all future requests are automatically queued, until the problem is solved and this mode is manually deactivated."
		msg = "From: Blockchain-Service <" + emailSender + ">\n" + "To: " + to + "\n" + "Subject: [Blockchain-Service]: Security Alert\n\n" + body
	} else if msg == "Empty Account" {
		body := "This Ethereum account " + address + ", no longer has sufficient funds to be able to write other diplomas on the blockchain.\nPlease add more Ethereum to the account.\nAnother account has just taken over writing on the blockchain, if this is not the case the diplomas are put in queue."
		msg = "From:  Blockchain-Service <" + emailSender + ">\n" + "To: " + to + "\n" + "Subject: [Blockchain-Service]: Empty Account\n\n" + body
	}
	err := smtp.SendMail("smtp.gmail.com:587", auth, emailSender, []string{to}, []byte(msg))
	if err != nil {
		LogsError(err)
		return false
	}
	return true
}
