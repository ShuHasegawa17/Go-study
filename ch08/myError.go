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
}

func (se StatusErr) Error() string {
	return se.Message
}

func GenerateError(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}
