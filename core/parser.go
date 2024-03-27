package core

import (
	"encoding/json"
	"os"
)

type Generic interface {
	GenericType | Dependencies | ProjectTypes | Languages | PackagingType | JavaVserions | BootVersions
}

type GenericType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
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

func GetDependencies() []GenericType {
	var data Dependencies
	parseRegistriesFile(&data)

	dependencies := []GenericType{}
	for _, value := range data.Dependencies.Values {
		dependencies = append(dependencies, value.Dependencies...)
	}

	return dependencies

}

type Languages struct {
	Languages struct {
		Values []GenericType `json:"values"`
	} `json:"language"`
}

type PackagingType struct {
	Packaging struct {
		Values []GenericType `json:"values"`
	} `json:"packaging"`
}

func GetProjectTypes() []GenericType {
	var data ProjectTypes
	parseRegistriesFile(&data)
	return data.ProjectTypes.Types
}

type JavaVserions struct {
	Versions struct {
		Values []GenericType `json:"values"`
	} `json:"javaVersion"`
}

type BootVersions struct {
	Versions struct {
		Values []GenericType `json:"values"`
	} `json:"bootVersion"`
}

func GetBootVersions() []GenericType {
	var data BootVersions
	parseRegistriesFile(&data)
	return data.Versions.Values
}
func GetJavaVersions() []GenericType {
	var data JavaVserions
	parseRegistriesFile(&data)
	return data.Versions.Values
}
func GetLanguages() []GenericType {
	var data Languages
	parseRegistriesFile(&data)
	return data.Languages.Values
}

func GetPackagingTypes() []GenericType {
	var data PackagingType
	parseRegistriesFile(&data)
	return data.Packaging.Values
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
