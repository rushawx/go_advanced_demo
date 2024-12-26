package verify

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/request"
	"3-validation-api/pkg/response"
	"fmt"
	"net/http"
	"net/smtp"
	"net/textproto"
	"strconv"

	"github.com/google/uuid"
	"github.com/jordan-wright/email"
)

type VerifyHandlerDeps struct {
	*configs.Config
}

type VerifyHandler struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Send")

		body, err := request.HandleBody[SendRequest](&w, r)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		verificationUUID := uuid.New().String()
		fmt.Println("Verification UUID:", verificationUUID)

		e := email.Email{
			To:      []string{body.Email},
			From:    handler.Config.Email.Username,
			Subject: "Verify your email",
			Text:    []byte("Click the link to verify your email"),
			HTML:    []byte("<a href='http://localhost:8080/verify/" + verificationUUID + "'>Click here</a>"),
			Headers: textproto.MIMEHeader{},
		}
		e.Headers.Set("Content-Type", "text/html; charset=UTF-8")
		e.Send(
			handler.Config.Email.Host+":"+strconv.Itoa(handler.Config.Email.Port),
			smtp.PlainAuth(
				"",
				handler.Config.Email.Username,
				handler.Config.Email.Password,
				handler.Config.Email.Host,
			),
		)
		fmt.Println("Email sent to " + body.Email)

		data := map[string]string{"message": "Email sent to " + body.Email}
		response.Json(w, data, http.StatusOK)

		response.ToFile(body.Email, verificationUUID)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verify")

		hash := r.PathValue("hash")
		if hash == "" {
			http.Error(w, "Missing hash parameter", http.StatusBadRequest)
			return
		}

		body, err := request.HandleBody[VerifyRequest](&w, r)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		if !response.FromFile(body.Email, hash) {
			http.Error(w, "Invalid hash", http.StatusBadRequest)
			return
		}

		data := map[string]string{"message": "Email verified"}
		response.Json(w, data, http.StatusOK)
	}
}
