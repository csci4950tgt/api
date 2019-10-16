package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	_, err := http.Get("http://127.0.0.1:8080/api")
	if err != nil {
		t.Error(err)
	}
}

func TestCreateClient(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	url := "http://127.0.0.1:8080/api/honeyclient/create"
	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	statusCode := resp.StatusCode
	if statusCode != http.StatusMethodNotAllowed {
		t.Error("Expected http status code: 405, actual code: " + resp.Status)
	}

	var requestJson = []byte(`{"url": "https://www.google.com","screenshot": {"width": 1920,"height": 1080,"filename": "screenshot.png"},"useragent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1 Safari/605.1.15"}`)
	resp, err = http.Post(url, "application/json", bytes.NewBuffer(requestJson))

	statusCode = resp.StatusCode
	if statusCode != http.StatusOK {
		t.Error("HTTP error, reason " + resp.Status)
	}
}

func TestGetHoneyClientById(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	resp, err := http.Get("http://127.0.0.1:8080/api/honeyclient/5")
	if err != nil {
		t.Error(err)
	}

	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		t.Error("HTTP error, reason " + resp.Status)
	}
}
