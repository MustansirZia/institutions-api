package providers

import "path/filepath"

func NewWorldUniversitiesProvider() InstitutionProvider {
	return NewJSONProvider(filepath.Join("data", "json", "world_universities_and_domains.json"), func(institution map[string]interface{}) string {
		return institution["name"].(string)
	})
}
