package err_handler

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
)

var (
	notFoundErrors = []error{
		sql.ErrNoRows,
	}
	badRequestErrors = []error{}
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var ferr *fiber.Error
	if errors.As(err, &ferr) {
		ctx.Status(ferr.Code)
	}

	if isNotFoundError(err) {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"data":  "",
			"error": err.Error(),
		})
	}
	if isBadRequest(err) {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"data":  "",
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data":  "",
		"error": err.Error(),
	})

}

func isNotFoundError(err error) bool {
	for _, e := range notFoundErrors {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}

func isBadRequest(err error) bool {
	for _, e := range badRequestErrors {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}
