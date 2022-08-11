package models

import (
	"context"
	"database/sql"
	"log"
	r "math/rand"
	"os"
	"reflect"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateMySQLConnection() *sql.DB {
	log.Println("Starting connetion to mysql...")
	os.Setenv("DSN", "root:password@tcp(localhost:3306)/mysql")
	DSN := os.Getenv("DSN")
	retries := 0
	for {
		if retries < 10 {
			db, err := sql.Open("mysql", DSN)
			if err != nil {
				log.Println("Connection problem, retrying ...")
				log.Println(err)
				retries++
				time.Sleep(time.Second * 5)
			} else {
				log.Println("Connected successfully")
				return db
			}
		} else {
			log.Panic("Connection failed!")
			return nil
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
func TestOrder_Save(t *testing.T) {
	db := CreateMySQLConnection()
	SetDb(db)
	EnsureTableExist(db)
	id := r.Intn(10000000)
	tests := []struct {
		name    string
		o       *Order
		wantErr bool
	}{
		{name: "should Saves successfull", o: &Order{OrderId: id, Price: 200, Title: "Burger"}, wantErr: false},
		{name: "should gets duplicate error", o: &Order{OrderId: id, Price: 200, Title: "Burger"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Order.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewOrder(t *testing.T) {
	tests := []struct {
		name string
		want IOrder
	}{
		{name: "type should be a pointer to models.Order ", want: NewOrder()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
