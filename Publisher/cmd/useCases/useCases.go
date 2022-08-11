package usecases

import (
	"log"

	"github.com/nelthaarion/snappfood-test-project-publisher/cmd/repositories"
)

type IUseCase interface {
	SendToRedis() error
	SetOrder(order repositories.Order) IUseCase
}

type UseCase struct {
	order repositories.Order
}

func (u *UseCase) SendToRedis() error {
	//logging or more needed logics goes here
	if err := u.order.SendToRedis(); err != nil {
		log.Println("somthing went wrong")
		log.Println(err)
		return err
	}
	return nil
}
func (u *UseCase) SetOrder(order repositories.Order) IUseCase {
	u.order = order
	return u
}

func NewUseCase() IUseCase {
	return &UseCase{}
}
