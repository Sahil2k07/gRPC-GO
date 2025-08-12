package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/Sahil2k07/gRPC-GO/internal/config"
	"github.com/Sahil2k07/gRPC-GO/internal/enum"
	errz "github.com/Sahil2k07/gRPC-GO/internal/error"
	"github.com/Sahil2k07/gRPC-GO/internal/model"

	"slices"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserData struct {
	ID    uint
	Email string
	Roles []enum.Role
}

func GenerateJWT(m model.User) (string, error) {
	JwtConfig := config.GetJWTConfig()

	claims := jwt.MapClaims{
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
		"id":    m.ID,
		"email": m.Email,
		"roles": m.Roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JwtConfig.Secret))
}

func GetUserFromToken(c echo.Context) (*UserData, error) {
	user := c.Get("user")
	if user == nil {
		return nil, errz.NewUnauthorized("user claims not found")
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		return nil, errz.NewUnauthorized("invalid token type")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errz.NewUnauthorized("invalid claims type")
	}

	// "id" is float64 in claims, convert to uint
	idFloat, ok := claims["id"].(float64)
	if !ok {
		return nil, errz.NewUnauthorized("id claim missing or invalid")
	}
	id := uint(idFloat)

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errz.NewUnauthorized("email claim missing or invalid")
	}

	roles, ok := claims["roles"].(string)
	if !ok {
		return nil, errz.NewUnauthorized("roles claim missing or invalid")
	}

	return &UserData{
		ID:    id,
		Email: email,
		Roles: parseRoles(roles),
	}, nil
}

func parseRoles(rolesStr string) []enum.Role {
	parts := strings.Split(rolesStr, ",")
	roles := make([]enum.Role, 0, len(parts))
	for _, p := range parts {
		roles = append(roles, enum.Role(p))
	}
	return roles
}

func WithRole(allowedRoles ...enum.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userData, err := GetUserFromToken(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			if slices.Contains(userData.Roles, enum.ADMIN) {
				return next(c)
			}

			for _, userRole := range userData.Roles {
				if slices.Contains(allowedRoles, userRole) {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "forbidden: insufficient permissions")
		}
	}
}
