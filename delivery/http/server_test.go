package http

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"tokeno/domain"
	"tokeno/usecase"
	"tokeno/utils"
)

func TestReturnBadRequest(t *testing.T) {
	type args struct {
		res          *httptest.ResponseRecorder
		responseData domain.FailedResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test bad request", args: struct {
			res          *httptest.ResponseRecorder
			responseData domain.FailedResponse
		}{res: httptest.NewRecorder(), responseData: domain.FailedResponse{Message: "gagal"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReturnBadRequest(tt.args.res, tt.args.responseData)
		})
		res := tt.args.res.Result()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Http status code should be bad request")
		}
		data, _ := io.ReadAll(res.Body)
		obj := utils.FromJson[domain.FailedResponse](string(data))
		if obj.Message != "gagal" {
			t.Errorf("Http message doesnt match")
		}
		if obj.Code != fmt.Sprintf("%d", http.StatusBadRequest) {
			t.Errorf("body status code should be bad request")
		}
	}
}

func TestReturnNotFoundResponse(t *testing.T) {
	type args struct {
		res          *httptest.ResponseRecorder
		responseData domain.FailedResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test not found", args: struct {
			res          *httptest.ResponseRecorder
			responseData domain.FailedResponse
		}{res: httptest.NewRecorder(), responseData: domain.FailedResponse{Message: "gagal"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReturnNotFoundResponse(tt.args.res, tt.args.responseData)
		})
		res := tt.args.res.Result()
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("Http status code should be bad request")
		}
		data, _ := io.ReadAll(res.Body)
		obj := utils.FromJson[domain.FailedResponse](string(data))
		if obj.Message != "gagal" {
			t.Errorf("Http message doesnt match")
		}
		if obj.Code != fmt.Sprintf("%d", http.StatusNotFound) {
			t.Errorf("body status code should be bad request")
		}
	}
}

func TestReturnSuccess(t *testing.T) {
	type args struct {
		res          *httptest.ResponseRecorder
		responseData domain.SuccessResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test searching id 1 success", args: struct {
			res          *httptest.ResponseRecorder
			responseData domain.SuccessResponse
		}{res: httptest.NewRecorder(), responseData: domain.SuccessResponse{Data: []domain.Data{
			{Id: 1, Name: "A"},
		}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReturnSuccess(tt.args.res, tt.args.responseData)
		})
		res := tt.args.res.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("Http status code should be ok")
		}
		data, _ := io.ReadAll(res.Body)
		obj := utils.FromJson[domain.SuccessResponse](string(data))
		dataList := obj.Data.([]any)
		if len(dataList) != 1 {
			t.Errorf("data length should be 1")
			dataContent := dataList[0].(map[string]any)
			if dataContent["id"] != 1 {
				t.Errorf("data content id should be 1")
			}
			if dataContent["name"] != "A" {
				t.Errorf("data content name should be A")
			}
		}
		if obj.Code != fmt.Sprintf("%d", http.StatusOK) {
			t.Errorf("body status code should be ok")
		}
	}
}

func Test_httpServer_handleRequest_badRequest(t *testing.T) {
	tokenoUseCase := usecase.NewTokenoUseCase()
	tokenoUseCase.InitData()
	type fields struct {
		usecase domain.TokenoUseCase
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "return bad request", fields: fields{usecase: tokenoUseCase}, args: struct {
			w *httptest.ResponseRecorder
			r *http.Request
		}{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/?id=dsa", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				usecase: tt.fields.usecase,
			}
			s.handleRequest(tt.args.w, tt.args.r)
		})
		res := tt.args.w.Result()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Http status code should be bad request")
		}
		data, _ := io.ReadAll(res.Body)
		obj := utils.FromJson[domain.FailedResponse](string(data))
		if obj.Message != "Invalid or Empty ID : [dsa]" {
			t.Errorf("Http message doesnt match, should be %s, but return %s", "Invalid or Empty ID : [dsa]", obj.Message)
		}
		if obj.Code != fmt.Sprintf("%d", http.StatusBadRequest) {
			t.Errorf("body status code should be bad request")
		}
	}
}

func Test_httpServer_handleRequest_notFound(t *testing.T) {
	tokenoUseCase := usecase.NewTokenoUseCase()
	tokenoUseCase.InitData()
	type fields struct {
		usecase domain.TokenoUseCase
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "return bad request", fields: fields{usecase: tokenoUseCase}, args: struct {
			w *httptest.ResponseRecorder
			r *http.Request
		}{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/?id=4", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				usecase: tt.fields.usecase,
			}
			s.handleRequest(tt.args.w, tt.args.r)
		})
		res := tt.args.w.Result()
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("Http status code should be not found")
		}
		data, _ := io.ReadAll(res.Body)
		obj := utils.FromJson[domain.FailedResponse](string(data))
		if obj.Message != "resource with ID 4 not exist" {
			t.Errorf("Http message doesnt match, should be %s, but return %s", "resource with ID 4 not exist", obj.Message)
		}
		if obj.Code != fmt.Sprintf("%d", http.StatusNotFound) {
			t.Errorf("body status code should be not found")
		}
	}
}

func Test_httpServer_handleRequest_success(t *testing.T) {
	tokenoUseCase := usecase.NewTokenoUseCase()
	tokenoUseCase.InitData()
	type fields struct {
		usecase domain.TokenoUseCase
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "return bad request", fields: fields{usecase: tokenoUseCase}, args: struct {
			w *httptest.ResponseRecorder
			r *http.Request
		}{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/?id=1,2", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				usecase: tt.fields.usecase,
			}
			s.handleRequest(tt.args.w, tt.args.r)
		})
		res := tt.args.w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("Http status code should be ok")
		}
		data, _ := io.ReadAll(res.Body)
		obj := utils.FromJson[domain.SuccessResponse](string(data))
		dataList := obj.Data.([]any)
		// just test the size correct
		if len(dataList) != 2 {
			t.Errorf("return data length shold be 2")
		}
		if obj.Code != fmt.Sprintf("%d", http.StatusOK) {
			t.Errorf("body status code should be ok")
		}
	}
}
