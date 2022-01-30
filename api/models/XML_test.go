package models

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Custom type that allows setting the func that our Mock Do func will run instead
type MockDoType func(req *http.Request) (*http.Response, error) // MockClient is the mock client

type MockClient struct {
	MockDo MockDoType
}

// Overriding what the Do function should "do" in our MockClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestMockNoError(t *testing.T) {
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"XMLName": {   "Space": "",   "Local": "Data"  },  "ID": "1",  "Name": "MWNZ",  "Description": "..is awesome" }`))),
			}, nil
		},
	}

	_, err := XMLGet("https://test.com")

	if err != nil {
		t.Error("Unexpected error TestMockNoError")
		return
	}
	Client = &http.Client{}
}

func TestMockError(t *testing.T) {
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{}, errors.New("test")
		},
	}

	_, err := XMLGet("https://test.com")
	if err == nil {
		t.Error("expected error received nil")
		return
	}
	Client = &http.Client{}
}

func TestMockNon200(t *testing.T) {
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"XMLName": {   "Space": "",   "Local": "Data"  },  "ID": "1",  "Name": "MWNZ",  "Description": "..is awesome" }`))),
			}, nil
		},
	}

	_, err := XMLGet("https://test.com")

	if err == nil {
		t.Error("expected error received nil")
		return
	}

	Client = &http.Client{}
}

func TestXMLGetSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `response from the mock server goes here`)
	}))
	defer ts.Close()

	mockServerURL := ts.URL

	_, err := XMLGet(mockServerURL)
	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}
}
