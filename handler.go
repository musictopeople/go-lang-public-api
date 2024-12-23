package main

import (
	"fmt"
	"io"
	"net/http"
)

func (app *application) GetHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	data, err := app.fetchData(id)

	if err != nil {
		app.logger.Error(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

func (app *application) fetchData(id string) (string, error) {
	baseUrl := "https://jsonplaceholder.typicode.com/todos"
	if id != "" {
		baseUrl = "https://jsonplaceholder.typicode.com/todos/" + id
	}

	app.logger.Info("fetching data from " + baseUrl)

	response, err := app.client.Get(baseUrl)
	if err != nil {
		return "", fmt.Errorf("error making external request: %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(body), nil
}
