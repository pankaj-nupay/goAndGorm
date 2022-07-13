package models

import "gorm.io/gorm"

type Book struct {
	BookId    uint    `json:"bookid"`
	Bookname  string  `json:"bookName"`
	Bookdesc  string  `json:"bookDesc"`
	Bookgenre string  `json:"bookGenre"`
	Author    *Author `json:"author"`
}

type Author struct {
	Authorid   int    `json:"authorid"`
	AuthorName string `json:"authorname"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	return err
}
