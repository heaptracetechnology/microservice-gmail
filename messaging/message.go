package messaging

import (
	"encoding/json"
	"fmt"
	result "github.com/heaptracetechnology/microservice-gmail/result"
	"net/http"
	"net/smtp"
	"os"
)

type Email struct {
	Subject  string `json:"subject,omitempty"`
	Body     string `json:"message,omitempty"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Password string `json:"password,omitempty"`
}

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

//Send Email
func SendEmail(responseWriter http.ResponseWriter, request *http.Request) {

	var smtpHost = os.Getenv("SMTP_HOST")
	var smtpPort = os.Getenv("SMTP_PORT")

	decoder := json.NewDecoder(request.Body)

	var param Email
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	from := param.From
	to := param.To
	sub := param.Subject
	body := param.Body
	smtpAddress := smtpHost + ":" + smtpPort

	msg := "From: " + from + "\n" + "To: " + to + "\n" + "Subject: " + sub + "\n" + body

	err := smtp.SendMail(smtpAddress, smtp.PlainAuth("", from, param.Password, smtpHost), from, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("err ::", err)
		return
	} else {
		message := Message{"true", "Email sent", http.StatusOK}
		bytes, _ := json.Marshal(message)
		result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
	}

}
