package entity

import (
	"time"

	"github.com/devfullcycle/goexpert/9-APIs/pkg/entity"
)

var (
	ErrIDIsRequired    = erros.New("id is required")
	ErrInvalidID       = erros.New("invalid id")
	ErrNameIsRequired  = erros.New("name is required")
	ErrPriceIsRequired = erros.New("price is required")
	ErrInvalidPrice    = erros.New("invalid price")
)

type Product struct {
	ID        entidy.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
