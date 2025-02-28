package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func newReverseProxy(target string) *httputil.ReverseProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatal("Ошибка при разборе URL:", err)
	}
	return httputil.NewSingleHostReverseProxy(targetURL)
}

func main() {
	authServiceURL := "http://auth_service:8081" // auth_service
	blogServiceURL := "http://blog_service:8082" // blog_service

	authProxy := newReverseProxy(authServiceURL)
	blogProxy := newReverseProxy(blogServiceURL)

	http.HandleFunc("/auth/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/auth")
		authProxy.ServeHTTP(w, r)
	})

	http.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/blog")
		blogProxy.ServeHTTP(w, r)
	})

	log.Println("Gateway-сервис запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
