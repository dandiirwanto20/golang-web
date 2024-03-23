package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

// implement unit test for http web golang
func TestHttp(t *testing.T) {
	// test request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	// test response
	recorder := httptest.NewRecorder()

	// initialize HelloHandler func
	HelloHandler(recorder, request)

	// check result test
	response := recorder.Result()
	// implement for can read all the response, in here to get body response
	body, _ := io.ReadAll(response.Body)
	// convertion []byte to string result
	bodyString := string(body)

	fmt.Println(bodyString)
}
