package providers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type NameExtractor func(data map[string]interface{}) string

type jsonProvider struct {
	jsonPath  string
	extractor NameExtractor
}

func NewJSONProvider(jsonPath string,
	extractor NameExtractor) InstitutionProvider {
	return &jsonProvider{
		jsonPath:  jsonPath,
		extractor: extractor,
	}
}

func (p *jsonProvider) Provide() ([]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(wd, p.jsonPath))
	if err != nil {
		return nil, err
	}
	data := make([]map[string]interface{}, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	institutions := make([]string, 0, len(data))
	for _, institution := range data {
		institutions = append(institutions, p.extractor(institution))
	}
	return institutions, nil
}
