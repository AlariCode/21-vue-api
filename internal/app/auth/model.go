package auth

// LoginRequest представляет запрос для аутентификации
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse представляет ответ после успешной аутентификации
type LoginResponse struct {
	Token string `json:"token"`
}

// ProfileResponse представляет данные профиля пользователя
type ProfileResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
