package main

import (
	"tokeno/delivery/http"
	"tokeno/usecase"
)

func main() {
	usecase := usecase.NewTokenoUseCase()
	usecase.InitData()
	http.NewHttpServer(3000, usecase)
}
