package main

import (
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreateMySQLConnection(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "should return pointer to sql.DB",
			args:    args{dsn: os.Getenv("DSN")},
			wantErr: false,
		},
		{
			name:    "should return nil with error",
			args:    args{dsn: "wrongdsn"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.args.dsn, tt.name)
			_, err := CreateMySQLConnection(tt.args.dsn)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMySQLConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCreateRedisClient(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "should return pointer to redis.Client",
			args:    args{address: os.Getenv("REDIS_ADDRESS")},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.args.address, tt.name)
			_, err := CreateRedisClient(tt.args.address)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRedisClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
