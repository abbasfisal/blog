package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "The field is required",
		"email":    "The field must have valid email",
		"min":      "Should be more than limit",
		"max":      "Should be less than limit",
	}
}
