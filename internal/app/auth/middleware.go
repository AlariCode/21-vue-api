package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWTMiddleware проверяет валидность JWT токена
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Получение токена из заголовка Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Отсутствует токен авторизации",
			})
		}

		// Проверка формата токена "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Неверный формат токена",
			})
		}
		tokenString := tokenParts[1]

		// Парсинг и валидация токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Проверка алгоритма подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Неверный метод подписи токена")
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Недействительный токен",
			})
		}

		// Токен валиден, извлекаем данные из claims и добавляем в контекст
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("email", claims["email"])
		}

		// Продолжаем выполнение запроса
		return c.Next()
	}
}
