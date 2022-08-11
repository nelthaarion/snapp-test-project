package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nelthaarion/snappfood-test-project-consumer/cmd/models"
	usecases "github.com/nelthaarion/snappfood-test-project-consumer/cmd/useCases"
)

var (
	DSN           = os.Getenv("DSN")
	REDIS_ADDRESS = os.Getenv("REDIS_ADDRESS")
)

func main() {
	client, err := CreateRedisClient(REDIS_ADDRESS)
	if err != nil {
		log.Panic("connection failed!")
	}
	sqlClient, err := CreateMySQLConnection(DSN)
	if err != nil {
		log.Panic("connection failed!")
	}
	EnsureTableExist(sqlClient)
	models.SetDb(sqlClient)
	for {
		result, _ := client.BLPop(context.TODO(), 0*time.Second, "ordersQueue").Result()
		u := usecases.NewUseCase(result[1])
		if err := u.Save(); err != nil {
			log.Println(err)
		} else {
			log.Println("order saved")
		}
	}
}
