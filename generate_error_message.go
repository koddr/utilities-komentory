package utilities

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GenerateErrorMessage func for generating error messages.
//
// 400 -> incorrect data for the <object> (<explanation>)
// 401 -> unauthorized request to the <object> (<explanation>)
// 403 -> access denied to the <object> (<explanation>)
// 404 -> <object> not found (<explanation>)
// 500 -> something went wrong with the <object> (<explanation>)
// XXX -> oops... something went wrong, but the status code is not handled (<code>)
func GenerateErrorMessage(statusCode int, object, explanation string) string {
	var errMsg string
	switch statusCode {
	case fiber.StatusBadRequest:
		errMsg = fmt.Sprintf("incorrect data for the %s (%s)", object, explanation)
	case fiber.StatusUnauthorized:
		errMsg = fmt.Sprintf("unauthorized request to the %s (%s)", object, explanation)
	case fiber.StatusForbidden:
		errMsg = fmt.Sprintf("access denied to the %s (%s)", object, explanation)
	case fiber.StatusNotFound:
		errMsg = fmt.Sprintf("%s not found (%s)", object, explanation)
	case fiber.StatusInternalServerError:
		errMsg = fmt.Sprintf("something went wrong with the %s (%s)", object, explanation)
	default:
		errMsg = fmt.Sprintf("oops... something went wrong, but the status code is not handled (%d)", statusCode)
	}
	return errMsg
}

// CheckForError func for checking error in given function and returning error message by status code.
func CheckForError(ctx *fiber.Ctx, errFunc error, statusCode int, object, explanation string) error {
	if errFunc != nil {
		return ctx.Status(statusCode).JSON(&fiber.Map{
			"error": true,
			"msg":   GenerateErrorMessage(statusCode, object, explanation),
		})
	}
	return nil
}
