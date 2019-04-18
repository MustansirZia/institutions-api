package providers

func NewWorldUniversitiesProvider() InstitutionProvider {
	return NewJSONProvider("./data/json/world_universities_and_domains.json", func(institution map[string]interface{}) string {
		return institution["name"].(string)
	})
}
