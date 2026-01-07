package utils

func GetValueOrDefault(value, defaultValue string) string {
	if value != "" {
		return value
	}
	return defaultValue
}