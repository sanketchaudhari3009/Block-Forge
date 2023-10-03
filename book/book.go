package book

import (
	"crypto/md5"
	"fmt"
	"io"
)

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}

func NewBook() *Book {
	return &Book{}
}

func (b *Book) GenerateID() {
	h := md5.New()
	io.WriteString(h, b.ISBN+b.PublishDate)
	b.ID = fmt.Sprintf("%x", h.Sum(nil))
}
