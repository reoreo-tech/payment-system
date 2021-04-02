package domain

// サーバで使うデータをまとめる

type Item struct {
	ID          int64
	Name        string
	Description string
	Amount      int64
}

// Items -set of item list
type Items []Item
