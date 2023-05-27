package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

func MapHandler(pathsToUrl map[string]string, mux http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newURL, ok := pathsToUrl[r.URL.Path]
		if ok {
			http.Redirect(w, r, newURL, http.StatusFound)
			return
		}
		mux.ServeHTTP(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/hello", hello)
	fmt.Println("Server running on port 8080")
	pathsToUrl := map[string]string{
		"/github-profile": "https://github.com/parthmalavia",
		"/go-hanson":      "https://courses.calhoun.io/courses/cor_gophercises",
	}
	urlMapHandler := MapHandler(pathsToUrl, mux)
	http.ListenAndServe(":8080", urlMapHandler)
	// s := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mux,
	// }
	// s.ListenAndServe()
}
