package main

import (
	"io/ioutil"
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
	resp, err := http.Get("http://127.0.0.1:8080/api/create-client")
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	response := string(body)
	if len(response) == 0 {
		t.Error("No response received")
	}
}

func TestGetHoneyClientById(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	resp, err := http.Get("http://127.0.0.1:8080/api/5")
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	response := string(body)
	if len(response) == 0 {
		t.Error("No response received")
	}
}
