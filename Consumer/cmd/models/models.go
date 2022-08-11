package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

var (
	dbTimeout = time.Second * 5
	client    *sql.DB
)

type IOrder interface {
	Save() error
}

type Order struct {
	OrderId int    `json:"order_id"`
	Price   int    `json:"price"`
	Title   string `json:"title"`
}

func (o *Order) Save() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	query := fmt.Sprintf(`INSERT INTO orders(order_id,price,title) VALUES (%d,%d,"%s")`, o.OrderId, o.Price, o.Title)
	defer cancel()
	_, err := client.QueryContext(ctx, query)
	return err
}

func NewOrder() IOrder {
	return &Order{}
}

func SetDb(db *sql.DB) {
	client = db
}
