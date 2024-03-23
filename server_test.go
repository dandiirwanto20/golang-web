package golangweb

import (
	"net/http"
	"testing"
)

// case make server
func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	// handle error in go
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
