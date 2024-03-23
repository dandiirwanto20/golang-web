package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	// writer is response and request is request from client
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		// menggunakan format function untuk auto convert type data
		fmt.Fprint(writer, "Hello World") // untuk menampilkan content di web, alternatif dari writer.Write()
	}

	// make server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	// checking error
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// implement ServeMux alternative for use multiple handler (not only in domain (/))
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	// implement multiple handler (enpoint)
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	// implement multiple handler (endpoint)
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Dandi")
	})

	// implement url pattern
	// url pattern is long endpoint is priority
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Image")
	})
	mux.HandleFunc("/images/thumbnails", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Thumbnail")
	})

	// make server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	// check error
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintln(writer, request.Method)     // for checking the request method
		fmt.Fprintln(writer, request.RequestURI) // for checking the request URI / url
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
