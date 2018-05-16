package http

import (
	"net/http"
	"strings"

	"github.com/open-falcon/mail-provider/config"
	"github.com/toolkits/smtp"
	"github.com/toolkits/web/param"

	"gopkg.in/gomail.v2"
	"strconv"
)

func configProcRoutes() {

	http.HandleFunc("/sender/mail", func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Config()
		token := param.String(r, "token", "")
		if cfg.Http.Token != token {
			http.Error(w, "no privilege", http.StatusForbidden)
			return
		}

		tos := param.MustString(r, "tos")
		subject := param.MustString(r, "subject")
		content := param.MustString(r, "content")
		tos = strings.Replace(tos, ",", ";", -1)

		if cfg.Smtp.Type == "smtp_ssl" {
			m := gomail.NewMessage()
			m.SetHeader("From", cfg.Smtp.From)
			m.SetHeader("To", tos)
			//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
			m.SetHeader("Subject", subject)
			m.SetBody("text/html", content)
			//m.Attach("/home/Alex/lolcat.jpg")

			d := gomail.NewDialer(cfg.Smtp.Addr, cfg.Smtp.Port, cfg.Smtp.Username, cfg.Smtp.Password)

			// Send the email to Bob, Cora and Dan.
			if err := d.DialAndSend(m); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}else {
				http.Error(w, "success", http.StatusOK)
			}
		}else {
			s := smtp.New(cfg.Smtp.Addr+":" + strconv.Itoa(cfg.Smtp.Port), cfg.Smtp.Username, cfg.Smtp.Password)
			err := s.SendMail(cfg.Smtp.From, tos, subject, content)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				http.Error(w, "success", http.StatusOK)
			}
		}
	})

}
