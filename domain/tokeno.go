package domain

type Data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type BaseResponse struct {
	Code string `json:"Code"`
}

type SuccessResponse struct {
	BaseResponse
	Data any `json:"data"`
}

type FailedResponse struct {
	BaseResponse
	Message string `json:"Message"`
}

type TokenoUseCase interface {
	InitData()
	Add(data *Data) error
	Get(ids []int) []Data
}
