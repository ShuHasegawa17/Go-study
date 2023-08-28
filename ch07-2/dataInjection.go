package main

import (
	"errors"
	"fmt"
	"net/http"
)

// ログを記録する関数
func LogOutput(message string) {
	fmt.Println(message)
}

// アプリのデータを保存する場所(SimpleDataStore)
type SimpleDataStore struct {
	userData map[string]string
}
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name,ok
}

// SimpleDataStoreのインスタンスを作成するファクトリ関数
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "hase",
			"2": "taro",
			"3": "hana",
		},
	}
}

// ビジネスロジックが依存するインターフェース
type DataStore interface {
	UserNameForID(userID string)(string,bool)
}
type Logger interface {
	Log(message string)
}

// LogOutputが上記のインターフェースに適合するように関数型を定義
type LoggerAdapter func(message string)
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// ビジネスロジック
// 定義したフィールドはどちらも具象型（LogOutput, SimpleDataStore）とは依存していない
type SimpleLogic struct {
	l Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string,error) {
	sl.l.Log("SayHello(" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("不明なユーザ")
	}
	return name + "さんこんにちは", nil

}

func (sl SimpleLogic) SayGoodbye(userID string) (string,error) {
	sl.l.Log("SayGoodbye(" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("不明なユーザ")
	}
	return name + "さんさようなら", nil
}

// ビジネスロジックのファクトリ関数
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l: l,
		ds: ds,
	}
}

// ビジネスロジックから、コントローラーが利用するものを定義
type Logic interface {
	SayHello(userID string) (string, error)
}

// コントローラーの作成
type Controller struct {
	l Logger
	logic Logic
}
func (c Controller) HandleGreeting(w http.ResponseWriter, r *http.Request) {
	c.l.Log("SayHello内")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}
// コントローラのファクトリ関数
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l: l,
		logic: logic,
	}
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l,ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.HandleGreeting)
	http.ListenAndServe(":8080",nil)

}