package main


import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func hi (w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Hi\n")
}

func main() {

	//http.HandleFunc("/hello",  hello)
	//http.HandleFunc("/headers", headers)
	//http.HandleFunc("/", hi)
	//http.ListenAndServe(":8090", nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8090", nil))

}
