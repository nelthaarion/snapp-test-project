package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nelthaarion/snappfood-test-project-consumer/cmd/models"
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
func CreateMySQLConnection(dsn string) (*sql.DB, error) {
	log.Println("Starting connetion to mysql...")
	retries := 0
	for {
		if retries < 5 {
			db, err := sql.Open("mysql", dsn)
			log.Println(dsn, err, db)
			if err != nil {
				log.Println("Connection problem, retrying ...")
				log.Println(err)
				retries++
				time.Sleep(time.Second * 5)
			} else {
				log.Println("Connected successfully")
				return db, nil
			}
		} else {
			return nil, errors.New("connection failed")
		}
	}
}
func EnsureTableExist(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	query := `CREATE TABLE IF NOT EXISTS orders(
		order_id int NOT NULL auto_increment,
		price int NOT NULL ,
		title varchar(30) NOT NULL ,
		PRIMARY KEY(order_id)
	);`
	_, err := db.QueryContext(ctx, query)
	return err
}

type IResponseData interface {
	SetMessage(msg string) IResponseData
	SetData(json models.IOrder) IResponseData
	SetError(isThereError bool) IResponseData
	Send(res http.ResponseWriter, statusCode int)
}
type ResponseData struct {
	Message string        `json:"message"`
	Data    models.IOrder `json:"data"`
	Error   bool          `json:"error"`
}

func GenerateResponse() *ResponseData {
	return &ResponseData{}
}
func (r *ResponseData) SetMessage(msg string) IResponseData {
	r.Message = msg
	return r
}
func (r *ResponseData) SetData(data models.IOrder) IResponseData {
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
