package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey = "your-secret-key" // В реальном приложении ключ следует хранить в переменных окружения
)

// Login обрабатывает запрос на авторизацию
func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат запроса",
		})
	}

	// Проверка учетных данных (простая проверка)
	if req.Email != "a@a.ru" || req.Password != "1" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Неверный email или пароль",
		})
	}

	// Создание JWT токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": req.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Срок действия 24 часа
	})

	// Подписываем токен
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка создания токена",
		})
	}

	return c.JSON(LoginResponse{Token: tokenString})
}

// GetProfile возвращает данные профиля пользователя
func GetProfile(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	return c.JSON(ProfileResponse{
		Email: email,
		Name:  "Антон",
	})
}
