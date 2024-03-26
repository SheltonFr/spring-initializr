package core

import (
	"io"
	"net/http"
	"os"
)

const specsURL string = "https://start.spring.io/metadata/client"

func FetchRegistries() error {

	result, err := http.Get(specsURL)
	if err != nil {
		return err
	}

	defer result.Body.Close()

	file, err := os.Create("registries.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, result.Body)

	if err != nil {
		return err
	}
	return nil
}
