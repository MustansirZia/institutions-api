package providers

import (
	"path/filepath"
	"strings"
)

func NewIndianCollegesProvider() InstitutionProvider {
	return NewJSONProvider(filepath.Join("data", "json", "indian_colleges.json"), func(institution map[string]interface{}) string {
		return strings.Trim(
			strings.Split(
				institution["College Name"].(string),
				"(Id:",
			)[0],
			" ",
		)
	})
}
