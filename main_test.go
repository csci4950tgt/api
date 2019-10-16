package main

import (
	"bytes"
	"encoding/json"
	"github.com/csci4950tgt/api/models"
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

	request := &models.Ticket{URL: "https://www.google.com", Screenshot: models.ScreenShot{Width: 1920, Height: 1080, Filename: "screenshot.png"}}
	requestJson, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
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
