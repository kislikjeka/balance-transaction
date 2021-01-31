package presenter

import (
	"TransactionTest/internal/domain"
	"time"
)

type Balance struct {
	ID        domain.ID `json:"id"`
	AccountID domain.ID `json:"account_id"`
	Balance   float32   `json:"balance"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func NewBalanceFromDomain(b *domain.Balance) *Balance {
	return &Balance{
		ID:        b.ID,
		AccountID: b.AccountId,
		Balance:   b.Balance,
		CreatedAt: b.CreatedAt.Format(time.RFC3339),
		UpdatedAt: b.UpdatedAt.Format(time.RFC3339),
	}
}
