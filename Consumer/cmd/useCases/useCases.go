package usecases

import "github.com/nelthaarion/snappfood-test-project-consumer/cmd/repositories"

type IUseCase interface {
	ConvertOrder() []byte
	Save() error
}

type UseCase struct {
	Order string
}

func (u *UseCase) ConvertOrder() []byte {
	return []byte(u.Order)
}

func (u *UseCase) Save() error {
	repo, err := repositories.GetRepo(u.ConvertOrder())
	if err != nil {
		return err
	}
	return repo.Save()
}

func NewUseCase(data string) IUseCase {
	return &UseCase{Order: data}
}
