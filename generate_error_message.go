package utilities

import "fmt"

// GenerateErrorMessage func for generating error messages:
//
// 400 -> incorrect data for the <object> (<context>)
// 401 -> unauthorized request to the <object> (<context>)
// 403 -> access denied to the <object> (<context>)
// 404 -> no <object> found (<context>)
// 500 -> something went wrong with the <object> (<context>)
func GenerateErrorMessage(statusCode int, object, context string) string {
	var errorMessage string
	switch statusCode {
	case 400:
		errorMessage = fmt.Sprintf("incorrect data for the %s (%s)", object, context)
	case 401:
		errorMessage = fmt.Sprintf("unauthorized request to the %s (%s)", object, context)
	case 403:
		errorMessage = fmt.Sprintf("access denied to the %s (%s)", object, context)
	case 404:
		errorMessage = fmt.Sprintf("no %s found (%s)", object, context)
	case 500:
		errorMessage = fmt.Sprintf("something went wrong with the %s (%s)", object, context)
	default:
		errorMessage = fmt.Sprintf("ooops... unhandled status code (%d)", statusCode)
	}
	return errorMessage
}
