package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

// Book model
type Book struct {
	gorm.Model
	Title  string `gorm:"size:100;not null;unique" json:"title"`
	Author string `gorm:"not null" json:"author"`
	Year   string `gorm:"not null" json:"year"`
}

// Prepare strips book inputs of any white space
func (b *Book) Prepare() {
	b.Title = strings.TrimSpace(b.Title)
	b.Author = strings.TrimSpace(b.Author)
	b.Year = strings.TrimSpace(b.Year)
}

// Validate validates the inputs
func (b *Book) Validate() error {
	if b.Title == "" {
		return errors.New("Title is required")
	}
	if b.Author == "" {
		return errors.New("Author is required")
	}
	if b.Year == "" {
		return errors.New("Year is required")
	}
	return nil
}

// Save stores the book object in the DB
func (b *Book) Save(db *gorm.DB) (*Book, error) {
	var err error

	// Debug a single operation, show detailed log for this operation
	err = db.Debug().Create(&b).Error
	if err != nil {
		return &Book{}, err
	}
	return b, nil
}

// GetBook returns the book
func (b *Book) GetBook(db *gorm.DB) (*Book, error) {
	book := &Book{}
	if err := db.Debug().Table("books").Where("title = ?", b.Title).First(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// GetBooks returns all the books
func GetBooks(db *gorm.DB) (*[]Book, error) {
	books := []Book{}
	if err := db.Debug().Table("books").Find(&books).Error; err != nil {
		return &[]Book{}, err
	}
	return &books, nil
}

// GetBookByID returns a book by its ID
func GetBookByID(id int, db *gorm.DB) (*Book, error) {
	book := &Book{}
	if err := db.Debug().Table("books").Where("id = ?", id).First(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// UpdateBook updates the provided book in the DB
func (b *Book) UpdateBook(id int, db *gorm.DB) (*Book, error) {
	if err := db.Debug().Table("books").Where("id = ?", id).Updates(Book{
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
	}).Error; err != nil {
		return &Book{}, err
	}
	return b, nil
}

// DeleteBook deletes a book from the DB by given ID
func DeleteBook(id int, db *gorm.DB) error {
	// Only deletes the book temporarily. To delete permanently use Table("books").Unscoped()
	if err := db.Debug().Table("books").Where("id = ?", id).Delete(&Book{}).Error; err != nil {
		return err
	}
	return nil
}
