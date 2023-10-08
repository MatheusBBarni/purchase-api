package config

import (
	"encoding/json"
	"net/http"
)

type HttpClient struct {
	client *http.Client
	url    string
}

func NewHttpClient(client *http.Client, url string) *HttpClient {
	return &HttpClient{
		client: client,
		url:    url,
	}
}

func (httpClient HttpClient) Get(target interface{}) error {
	r, err := httpClient.client.Get(httpClient.url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
