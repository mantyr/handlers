package api

import (
	"io/ioutil"
	"net/http/httptest"
	"reflect"
	"testing"
)

type TestAPIData struct {
	Key   string `json:"key" xml:"key,attr"`
	Value string `json:"value" xml:"value,attr"`
}

func TestAPI(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// nolint:dupl
		t.Run("JSON", func(t *testing.T) {
			t.Run("Status", func(t *testing.T) {
				t.Run("Default status", func(t *testing.T) {
					testAPI(
						t,
						"json",
						TestAPIData{"key", "value"},
						[]int{},
						[]byte(`{"key":"key","value":"value"}`),
						200,
					)
				})
				t.Run("300", func(t *testing.T) {
					testAPI(
						t,
						"json",
						TestAPIData{"key", "value"},
						[]int{300},
						[]byte(`{"key":"key","value":"value"}`),
						300,
					)
				})
			})
		})
		// nolint:dupl
		t.Run("XML", func(t *testing.T) {
			t.Run("Status", func(t *testing.T) {
				t.Run("Default status", func(t *testing.T) {
					testAPI(
						t,
						"xml",
						TestAPIData{"key", "value"},
						[]int{},
						[]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<TestAPIData key=\"key\" value=\"value\"></TestAPIData>"), // nolint:lll
						200,
					)
				})
				t.Run("300", func(t *testing.T) {
					testAPI(
						t,
						"xml",
						TestAPIData{"key", "value"},
						[]int{300},
						[]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<TestAPIData key=\"key\" value=\"value\"></TestAPIData>"), // nolint:lll
						300,
					)
				})
			})
		})
	})
}

func testAPI(
	t *testing.T,
	responseType string,
	data interface{},
	status []int,
	expectedBody []byte,
	expectedStatus int,
) {
	resp := httptest.NewRecorder()
	api := New()
	var err error

	switch responseType {
	case "json":
		err = api.JSONResponse(resp, data, status...)
	case "xml":
		err = api.XMLResponse(resp, data, status...)
	}
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	response := resp.Result()
	if response == nil {
		t.Fatalf("empty response")
	}
	if response.StatusCode != expectedStatus {
		t.Fatalf(
			"expected %d status but actual %d status",
			expectedStatus,
			response.StatusCode,
		)
	}
	if response.Body == nil {
		t.Fatalf("empty body")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read body error: %v", err)
	}
	if !reflect.DeepEqual(body, expectedBody) {
		t.Fatalf(`expected %q but actual %q`, expectedBody, body)
	}
}
