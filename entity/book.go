package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Book struct {
	ID        ID
	Title     string
	Author    string
	Pages     int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBook(title string, author string, pages int, quantity int) (*Book, error) {
	b := &Book{
		ID:        NewID(),
		Title:     title,
		Author:    author,
		Pages:     pages,
		Quantity:  quantity,
		CreatedAt: time.Now(),
	}
	err := b.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return b, nil
}

func (b *Book) Validate() error {
	if b.Title == "" || b.Author == "" || b.Pages <= 0 || b.Quantity <= 0 {
		return ErrInvalidEntity
	}
	return nil
}

//AddBook add a book
func (u *User) AddBook(id ID) error {
	_, err := u.GetBook(id)
	if err == nil {
		return ErrBookAlreadyBorrowed
	}
	u.Books = append(u.Books, id)
	return nil
}

//RemoveBook remove a book
func (u *User) RemoveBook(id ID) error {
	for i, j := range u.Books {
		if j == id {
			// u.Books[:i] -> Pega tudo que tiver antes do indice, sem contar o que estiver no indice
			// u.Books[i+1:] -> Pega tudo apartir do indice
			u.Books = append(u.Books[:i], u.Books[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

//GetBook get a book
func (u *User) GetBook(id ID) (ID, error) {
	for _, v := range u.Books {
		if v == id {
			return id, nil
		}
	}
	return id, ErrNotFound
}

//ValidatePassword validate user password
func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}
