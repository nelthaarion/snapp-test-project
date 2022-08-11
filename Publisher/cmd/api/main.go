package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nelthaarion/snappfood-test-project-publisher/cmd/repositories"
)

func main() {
	redisClient, err := CreateRedisClient(os.Getenv("REDIS_ADDRESS"))
	if err != nil {
		log.Panic(err)
	}
	repositories.NewClient(redisClient)
	log.Println(os.Getenv("PORT"), os.Getenv("REDIS_ADDRESS"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), Routes()))
}
