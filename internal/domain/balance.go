package domain

import (
	"time"
)

type ID uint

type Balance struct {
	ID        ID
	AccountId ID
	Balance   float32
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewBalance create new Balance
func NewBalance(accountID ID) (*Balance, error) {
	b := &Balance{
		AccountId: accountID,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := b.Validate()
	if err != nil {
		return nil, err
	}

	return b, nil
}

//Validate check if Balance is valid
func (b *Balance) Validate() error {
	if b.AccountId < 0 {
		return ErrInvalidEntity
	}

	if b.Balance < 0 {
		return ErrInvalidBalance
	}

	return nil
}

func (b *Balance) ChangeBalance(val float32) error {
	if b.Balance+val < 0 {
		return ErrInvalidBalance
	}
	b.Balance = b.Balance + val

	return nil
}
