package registries

import (
	"errors"
	"io"
	"net/http"
	"os"
)

const url string = "https://start.spring.io/metadata/client"

func GenerateRegistriesFile() (string, error) {
	result, err := http.Get(url)

	if err != nil {
		return "", errors.New("an error occured while fetching registries")
	}

	defer result.Body.Close()

	file, err := os.Create("registries.json")
	if err != nil {
		return "", errors.New("an error occured while creating file")
	}
	defer file.Close()

	_, err = io.Copy(file, result.Body)

	if err != nil {
		panic(err)
	}

	return file.Name(), nil
}
