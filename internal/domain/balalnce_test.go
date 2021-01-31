package domain_test

import (
	"TransactionTest/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBalance(t *testing.T) {
	b, err := domain.NewBalance(123)
	assert.Nil(t, err)
	assert.Equal(t, b.AccountId, uint(123))
	assert.Equal(t, b.Balance, float32(0))
	assert.NotNil(t, b.CreatedAt)
	assert.NotNil(t, b.UpdatedAt)
}

func TestBalance_ChangeBalance(t *testing.T) {
	b, _ := domain.NewBalance(123)
	b.ChangeBalance(10)
	assert.Equal(t, float32(10), b.Balance)
	b.ChangeBalance(-5)
	assert.Equal(t, float32(5), b.Balance)
}
