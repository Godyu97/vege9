package vegeEmail

import (
	"strings"
	"net/smtp"
)

type EmailSendReq struct {
	User           string
	PassWord       string
	Host           string
	Subject        string
	Date           string
	Body           string
	MailType       string
	ReplyToAddress string
	To             []string
	CC             []string
	BCC            []string
}

func EmailSend(req EmailSendReq) error {
	hp := strings.Split(req.Host, ":")
	auth := smtp.PlainAuth("", req.User, req.PassWord, hp[0])
	var content_type string
	if req.MailType == "html" {
		content_type = "Content-Type: text/" + req.MailType + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	cc_address := strings.Join(req.CC, ";")
	bcc_address := strings.Join(req.BCC, ";")
	to_address := strings.Join(req.To, ";")
	msg := []byte("To: " + to_address + "\r\nFrom: " + req.User + "\r\nSubject: " + req.Subject + "\r\nDate: " + req.Date + "\r\nReply-To: " + req.ReplyToAddress + "\r\nCc: " + cc_address + "\r\nBcc: " + bcc_address + "\r\n" + content_type + "\r\n\r\n" + req.Body)
	send_to := MergeSlice(req.To, req.CC)
	send_to = MergeSlice(send_to, req.BCC)
	err := smtp.SendMail(req.Host, auth, req.User, send_to, msg)
	return err
}

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}
