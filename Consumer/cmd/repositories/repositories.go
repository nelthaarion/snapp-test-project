package repositories

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/nelthaarion/snappfood-test-project-consumer/cmd/models"
)

type IData interface {
	Save() error
}

type Data struct {
	Order models.IOrder
}

func (d *Data) Save() error {
	if d.Order != nil {
		log.Println(d.Order)
		return d.Order.Save()
	}
	return errors.New("order is empty")
}

func GetRepo(input []byte) (IData, error) {
	order := models.NewOrder()
	err := json.Unmarshal(input, &order)
	if err != nil {
		return nil, err
	}
	data := &Data{Order: order}
	return data, nil
}
