package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nelthaarion/snappfood-test-project-publisher/cmd/repositories"
)

var id = 1000

func TestHandleOrderWithoutError(t *testing.T) {
	client, _ := CreateRedisClient("localhost:6379")
	repositories.NewClient(client)
	payload := repositories.Order{OrderId: id, Price: 2000, Title: "Test burger"}
	data, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/order", strings.NewReader(string(data)))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	HandleOrder(res, req)
	response := res.Result()
	defer response.Body.Close()
	resData, _ := ioutil.ReadAll(response.Body)
	var responseMessage ResponseData
	json.Unmarshal(resData, &responseMessage)
	log.Println(responseMessage)
	if response.StatusCode != 201 || responseMessage.Error {
		t.Errorf("Test failed")
	}
}

func TestHandleOrderWithError(t *testing.T) {
	client, _ := CreateRedisClient("localhost:6379")
	repositories.NewClient(client)
	payload := "bad payload"
	data, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/order", bytes.NewBuffer(data))
	res := httptest.NewRecorder()
	HandleOrder(res, req)
	response := res.Result()
	defer response.Body.Close()
	resData, _ := ioutil.ReadAll(response.Body)
	var responseMessage ResponseData
	json.Unmarshal(resData, &responseMessage)
	log.Println(response.StatusCode, responseMessage.Error)
	if response.StatusCode == 201 || !responseMessage.Error {
		t.Errorf("Test failed")
	}
}
