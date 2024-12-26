package verify

import (
	"3-validation-api/configs"
	"encoding/json"
	"io"
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
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		var requestData struct {
			To string `json:"to"`
		}
		if err := json.Unmarshal(body, &requestData); err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		verificationUUID := uuid.New().String()

		e := email.Email{
			To:      []string{requestData.To},
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
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Email sent to " + requestData.To))
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.URL.Path[len("/verify/"):]
		if hash == "" {
			http.Error(w, "Missing hash parameter", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Verify " + hash))
	}
}
