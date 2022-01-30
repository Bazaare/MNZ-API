package models

import (
	"encoding/xml"
	"errors"
	"net/http"
)

var (
	Client HTTPClient
)

type XML struct {
	XMLName     xml.Name `xml:"Data"`
	ID          string   `xml:"id"`
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func init() {
	Client = &http.Client{}
}

func XMLGet(url string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
		//responses.ERROR(w, resp.StatusCode, errors.New(resp.Status))
	}

	return resp, nil
}
