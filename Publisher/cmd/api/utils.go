package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nelthaarion/snappfood-test-project-publisher/cmd/repositories"
)

func CreateRedisClient(address string) (*redis.Client, error) {
	log.Println("Starting connetion to redis...")
	retries := 0

	for {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "",
		})
		if retries < 5 {
			if redisClient == nil {
				log.Println("Retrying connection ...")
				retries++
				time.Sleep(time.Second * 2)
			} else {
				log.Println("Conneted successfully!")
				return redisClient, nil
			}
		} else {
			return nil, errors.New("connection failed")
		}

	}
}

type IResponseData interface {
	SetMessage(msg string) IResponseData
	SetData(json repositories.IOrder) IResponseData
	SetError(isThereError bool) IResponseData
	Send(res http.ResponseWriter, statusCode int)
}
type ResponseData struct {
	Message string              `json:"message"`
	Data    repositories.IOrder `json:"data"`
	Error   bool                `json:"error"`
}

func GenerateResponse() *ResponseData {
	return &ResponseData{}
}
func (r *ResponseData) SetMessage(msg string) IResponseData {
	r.Message = msg
	return r
}
func (r *ResponseData) SetData(data repositories.IOrder) IResponseData {
	r.Data = data
	return r
}
func (r *ResponseData) SetError(isThereError bool) IResponseData {
	r.Error = isThereError
	return r
}
func (r *ResponseData) Send(res http.ResponseWriter, statuscode int) {
	res.WriteHeader(statuscode)
	json.NewEncoder(res).Encode(r)
}
