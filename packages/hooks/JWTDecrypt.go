package hooks

import (
	"fmt"
	"strings"
	"time"

	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/conf"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID       string   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleID   uint   `json:"role_id"`
	Exp      int64  `json:"exp"`
}

func (c *Claims) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.Exp, 0)), nil
}

func (c *Claims) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (c *Claims) GetIssuedAt() (*jwt.NumericDate, error) {
	return nil, nil
}

func (c *Claims) GetIssuer() (string, error) {
	return "", nil
}

func (c *Claims) GetSubject() (string, error) {
	return "", nil
}

func (c *Claims) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

// Middleware to decrypt JWT
func DecryptJWT(c *fiber.Ctx) error {

	if c.Get("Authorization") == "" {
		return c.Next();
	}

	// Get token from the "Authorization" header
    authHeader := c.Get("Authorization")
    fmt.Println(authHeader);
    if authHeader == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
    }

    // Ensure token is in "Bearer <token>" format
    tokenParts := strings.Split(authHeader, " ")
    if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
    }

    tokenString := tokenParts[1]
    secretKey := conf.GetConfig().JWT.Secret

    // Parse the token
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secretKey), nil
    })

    if err != nil || !token.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
    }

    // Extract claims
    claims, ok := token.Claims.(*Claims)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
    }

    // Store user data in context for next handlers
    c.Locals("username", claims.Username)
	c.Locals("id", claims.ID)

    return c.Next()
}