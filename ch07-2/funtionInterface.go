package main

import "net/http"

type Hnaler interface {
	ServerHttp(http.ResponseWriter, *http.Request)
}

// 関数インターフェースの実装
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) 	ServerHttp(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
