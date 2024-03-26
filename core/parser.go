package core

import (
	"encoding/json"
	"os"
)

type Generic interface {
	GenericType | Dependencies | ProjectTypes
}

type GenericType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectTypes struct {
	ProjectTypes struct {
		Types []GenericType `json:"values"`
	} `json:"type"`
}

type Dependencies struct {
	Dependencies struct {
		Values []struct {
			Category     string        `json:"name"`
			Dependencies []GenericType `json:"values"`
		} `json:"values"`
	} `json:"dependencies"`
}

func GetProjectTypes() []GenericType {
	var data ProjectTypes
	parseRegistriesFile(&data)
	return data.ProjectTypes.Types
}

func parseRegistriesFile[T Generic](target *T) error {
	file, err := os.Open("registries.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(target)

	if err != nil {
		return err
	}
	return nil
}
