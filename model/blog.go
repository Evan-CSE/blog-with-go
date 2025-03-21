package model

type Blog struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (b *Blog) TableName() string {
	return "blog"
}
