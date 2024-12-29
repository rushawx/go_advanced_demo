package auth

type SessionRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type SessionResponse struct {
	SessionId string `json:"session_id"`
}

type LoginRequest struct {
	SessionId string `json:"session_id" validate:"required"`
	Code      string `json:"code" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}
