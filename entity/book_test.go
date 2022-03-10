package entity_test

import (
	"testing"

	"github.com/robsantossilva/clean-arch-go/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	b, err := entity.NewBook("Clean Architecture", "Uncle Bob", 400, 1)
	assert.Nil(t, err)
	assert.Equal(t, b.Title, "Clean Architecture")
	assert.NotNil(t, b.ID)
}

func TestBookValidate(t *testing.T) {
	type test struct {
		title    string
		author   string
		pages    int
		quantity int
		want     error
	}

	tests := []test{
		{
			title:    "American Gods",
			author:   "Neil Gaiman",
			pages:    100,
			quantity: 1,
			want:     nil,
		},
		{
			title:    "American Gods",
			author:   "Neil Gaiman",
			pages:    100,
			quantity: 0,
			want:     entity.ErrInvalidEntity,
		},
		{
			title:    "",
			author:   "Neil Gaiman",
			pages:    100,
			quantity: 1,
			want:     entity.ErrInvalidEntity,
		},
		{
			title:    "American Gods",
			author:   "",
			pages:    100,
			quantity: 1,
			want:     entity.ErrInvalidEntity,
		},
		{
			title:    "American Gods",
			author:   "Neil Gaiman",
			pages:    0,
			quantity: 1,
			want:     entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := entity.NewBook(tc.title, tc.author, tc.pages, tc.quantity)
		assert.Equal(t, err, tc.want)
	}

}
