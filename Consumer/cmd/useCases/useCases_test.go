package usecases

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nelthaarion/snappfood-test-project-consumer/cmd/models"
)

func CreateMySQLConnection() (*sql.DB, error) {
	log.Println("Starting connetion to mysql...")
	retries := 0
	for {
		if retries < 5 {
			db, err := sql.Open("mysql", os.Getenv("DSN"))

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
func TestUseCase_Save(t *testing.T) {
	db, _ := CreateMySQLConnection()
	models.SetDb(db)
	id := rand.Intn(100000000)
	data := fmt.Sprintf(`{"order_id":%d,"price":%d,"title":"%s"}`, id, 10, "Salad")
	log.Println(data)
	uc := NewUseCase(data)
	tests := []struct {
		name    string
		u       IUseCase
		wantErr bool
	}{
		{
			name:    "should save data without any error",
			u:       uc,
			wantErr: false,
		},
		{
			name:    "should get error for cannot marshal input data",
			u:       NewUseCase("bad data"),
			wantErr: true,
		},
		{
			name:    "should get error for duplicate data",
			u:       uc,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Save(); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
