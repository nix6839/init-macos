package utils

import (
	"io"
	"net/http"
)

func FetchBody(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
