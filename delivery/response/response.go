package response

import "net/mail"

func ResponseSuccess(message string, code int, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	}
}

func ResponseSuccessWithoutData(message string, code int) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": message,
	}
}

func ResponseFailed(message string, code int) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": message,
	}
}

func ValidMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}
