package providers

import (
	"path/filepath"
	"strings"
)

func NewIndianUniversitiesProvider() InstitutionProvider {
	return NewJSONProvider(filepath.Join("data", "json", "indian_universities.json"), func(institution map[string]interface{}) string {
		return strings.Trim(
			strings.Split(
				institution["University Name"].(string),
				"(Id:",
			)[0],
			" ",
		)
	})
}
