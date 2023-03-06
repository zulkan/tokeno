package usecase

import (
	"tokeno/domain"
)

type tokenoUseCase struct {
	mapData map[int]domain.Data
}

func (t *tokenoUseCase) InitData() {
	t.Add(&domain.Data{
		Id:   1,
		Name: "A",
	})
	t.Add(&domain.Data{
		Id:   2,
		Name: "B",
	})
	t.Add(&domain.Data{
		Id:   3,
		Name: "C",
	})
}

func (t *tokenoUseCase) Add(data *domain.Data) error {
	t.mapData[data.Id] = *data

	return nil
}

func (t *tokenoUseCase) Get(ids []int) []domain.Data {
	res := []domain.Data{}

	for _, val := range ids {
		if data, isExist := t.mapData[val]; isExist {
			res = append(res, data)
		}
	}

	return res
}

func NewTokenoUseCase() domain.TokenoUseCase {
	return &tokenoUseCase{
		mapData: map[int]domain.Data{},
	}
}
