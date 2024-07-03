package main

import "net/http"

func main() {

	http.HandleFunc("/hello", helloFunc)
	http.HandleFunc("/no", NoFunc)
	http.HandleFunc("/hi", hiio)
	add := "localhost:8080"
	err := http.ListenAndServe(add, nil)
	if err != nil {
		panic(err)
	}

}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "oops", http.StatusNotFound)
	}
	writerresponse(w, "hello..... we are!!!!!")
}

func NoFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "oops", http.StatusNotFound)
	}
	writerresponse(w, "no one here")

}
func hiio(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "oops", http.StatusNotFound)
	}
	writerresponse(w, "hiioo")
}
func writerresponse(w http.ResponseWriter, responseString string) {

	response := []byte(responseString)
	_, err := w.Write(response)
	if err != nil {
		panic(err)
	}
}
