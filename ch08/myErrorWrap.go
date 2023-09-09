package main

// ステータスの独自列挙型（iota)
type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

// 独自のエラー型定義
type StatusErr struct {
	Status  Status
	Message string
	Err     error // ラップするエラー
}

func (se StatusErr) Error() string {
	return se.Message
}

// アンラップを追加
func (se StatusErr) Unwrap() error {
	return se.Err
}
