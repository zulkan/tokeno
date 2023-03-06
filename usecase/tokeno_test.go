package usecase

import (
	"testing"
	"tokeno/domain"
	"tokeno/utils"
)

func TestAddAndGet(t *testing.T) {
	usecase := NewTokenoUseCase()
	usecase.InitData()
	res := usecase.Get([]int{1, 2})

	if len(res) != 2 {
		t.Errorf("result is not valid")

		if res[0].Id == res[1].Id {
			t.Errorf("data should be different")
		}
		if utils.InSlice[int]([]int{1, 2}, utils.MapSlice(res, func(data domain.Data) int {
			return data.Id
		})) {
			t.Errorf("result is not valid")
		}
	}

	res = usecase.Get([]int{4})
	if len(res) != 0 {
		t.Errorf("result is not valid")
	}
}
