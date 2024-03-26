package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	result, err := http.Get("https://start.spring.io/metadata/client")

	if err != nil {
		panic(err)
	}

	defer result.Body.Close()

	file, err := os.Create("registries.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, result.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Done!!!")
}
