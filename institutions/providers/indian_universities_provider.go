package providers

import (
	"strings"
)

func NewIndianUniversitiesProvider() InstitutionProvider {
	return NewJSONProvider("./data/json/indian_universities.json", func(institution map[string]interface{}) string {
		return strings.Trim(
			strings.Split(
				institution["University Name"].(string),
				"(Id:",
			)[0],
			" ",
		)
	})
}
