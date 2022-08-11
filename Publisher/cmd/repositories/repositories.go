package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

type IOrder interface {
	SendToRedis() error
	ToBytes() []byte
}
type Order struct {
	OrderId int    `json:"order_id"`
	Price   int    `json:"price"`
	Title   string `json:"title"`
}

func (o Order) SendToRedis() error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	_, err := client.RPush(ctx, "ordersQueue", o.ToBytes()).Result()
	return err
}
func (o Order) ToBytes() []byte {
	data, _ := json.Marshal(o)
	return data
}
func NewClient(c *redis.Client) {
	client = c
}
