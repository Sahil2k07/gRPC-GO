package util

import (
	"errors"
	"net/http"

	errz "github.com/Sahil2k07/gRPC-GO/internal/error"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HandleError(c echo.Context, err error) error {
	if err == nil {
		return nil
	}

	generateLogMessage(c, err)

	switch e := err.(type) {
	case *errz.NotFoundError:
		return c.JSON(http.StatusNotFound, map[string]string{"error": e.Error()})
	case *errz.ValidationError:
		return c.JSON(http.StatusBadRequest, map[string]string{"error": e.Error()})
	case *errz.UnauthorizedError:
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": e.Error()})
	case *errz.ForbiddenError:
		return c.JSON(http.StatusForbidden, map[string]string{"error": e.Error()})
	case *errz.AlreadyExistsError:
		return c.JSON(http.StatusConflict, map[string]string{"error": e.Error()})
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "resource not found"})
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
}
