package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func OkHandler(w http.ResponseWriter, r *http.Request) {
	expect := map[string]interface{}{"response": " passed "}
	ResponseOk(w, expect)
}

func ResponseOkHandler(w http.ResponseWriter, r *http.Request) {
	OkHandler(w, r)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	expect := "Error in nutshell"
	ResponseError(w, 500, expect)
}

func ResponseErrorHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r)
}

func TestResponseOk(t *testing.T) {
	t.Run("ResponseOk working properly", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/test", nil)

		if err != nil {
			t.Error(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ResponseOkHandler)
		handler.ServeHTTP(resp, req)

		if status := resp.Code; status != http.StatusOK {
			t.Errorf("Response for util is not ok  %v ", status)
		}

		expect := map[string]interface{}{"response": " passed "}
		var result map[string]interface{}

		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &result)

		if resp.Code != http.StatusOK {
			t.Errorf("Response for util is not ok")
		}

		if !reflect.DeepEqual(result, expect) {
			t.Errorf("Response for util is not ok  %v ", expect)
		}
	})
}

func TestResponseError(t *testing.T) {
	t.Run("Response Error working properly", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/test", nil)

		if err != nil {
			t.Error(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ResponseErrorHandler)
		handler.ServeHTTP(resp, req)

		expect := map[string]string{"error": "Error in nutshell"}
		var result map[string]string

		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &result)

		if resp.Code != http.StatusInternalServerError {
			t.Errorf("Response for util is not internal server error")
		}

		if !reflect.DeepEqual(result, expect) {
			t.Errorf("Response for util is not error  %v ", expect)
		}
	})
}
