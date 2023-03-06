package http

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"tokeno/domain"
	"tokeno/utils"
)

type httpServer struct {
	usecase domain.TokenoUseCase
}

func NewHttpServer(port int, usecase domain.TokenoUseCase) {
	server := httpServer{usecase: usecase}
	http.HandleFunc("/", server.handleRequest)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}

func (s *httpServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !isValidInput(query) {
		ReturnBadRequest(w, domain.FailedResponse{
			Message: fmt.Sprintf("Invalid or Empty ID : %v", query["id"]),
		})
		return
	}
	ids := strings.Split(query["id"][0], ",")
	idInts := utils.MapSlice[string, int](ids, func(data string) int {
		val, _ := strconv.ParseInt(data, 10, 64)
		return int(val)
	})
	res := s.usecase.Get(idInts)

	if len(res) == 0 {
		ReturnNotFoundResponse(w, domain.FailedResponse{
			Message: fmt.Sprintf("resource with ID %s not exist", strings.Join(ids, ",")),
		})
		return
	}
	ReturnSuccess(w, domain.SuccessResponse{
		Data: res,
	})
}

func isValidInput(query url.Values) bool {
	if _, exist := query["id"]; !exist || len(query["id"]) > 1 {
		return false
	}
	ids := strings.Split(query["id"][0], ",")
	if !utils.AllIsInt(ids) {
		return false
	}
	return true
}

func ReturnSuccess(res http.ResponseWriter, responseData domain.SuccessResponse) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	responseData.Code = fmt.Sprintf("%d", http.StatusOK)
	res.Write([]byte(utils.ToJson(responseData)))
}

func ReturnBadRequest(res http.ResponseWriter, responseData domain.FailedResponse) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusBadRequest)
	responseData.Code = fmt.Sprintf("%d", http.StatusBadRequest)
	res.Write([]byte(utils.ToJson(responseData)))
}

func ReturnNotFoundResponse(res http.ResponseWriter, responseData domain.FailedResponse) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusNotFound)
	responseData.Code = fmt.Sprintf("%d", http.StatusNotFound)
	res.Write([]byte(utils.ToJson(responseData)))
}
