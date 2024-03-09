package internal

import (
	"Convert/internal/models"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func Convert(input []byte) ([]byte, error) {
	person := models.Employee{}
	err := json.Unmarshal(input, &person)
	if err != nil {
		return nil, fmt.Errorf("ошибка при анмаршалинге из JSON: %w", err)
	}
	err = person.Validate()
	if err != nil {
		return nil, fmt.Errorf("данные не валидны: %w", err)
	}
	yamlData, err := yaml.Marshal(&person)
	if err != nil {
		return nil, fmt.Errorf("ошибка при маршалинге в YAML: %w", err)
	}
	return yamlData, nil
}
