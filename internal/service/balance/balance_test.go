package balance

import (
	"TransactionTest/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newFakeBalance() *domain.Balance {
	return &domain.Balance{
		AccountId: 4,
		Balance:   2.35,
	}
}

func Test_CreateAndGet(t *testing.T) {
	repo := newInmem()
	service := NewService(repo)
	bal := newFakeBalance()
	i, err := service.Create(bal.AccountId)
	assert.NotNil(t, i)
	assert.Nil(t, err)
	saved, err := service.Get(i)
	assert.Nil(t, err)
	assert.Equal(t, i, saved.ID)
	assert.Equal(t, float32(0), saved.Balance)

}

func Test_Update(t *testing.T) {
	repo := newInmem()
	service := NewService(repo)
	bal := newFakeBalance()
	i, _ := service.Create(bal.AccountId)
	saved, _ := service.Get(i)
	saved.Balance = 35
	assert.Nil(t, service.Update(saved))
	updated, _ := service.Get(i)
	assert.Equal(t, float32(35), updated.Balance)
}

func Test_List(t *testing.T) {
	repo := newInmem()
	service := NewService(repo)
	b1 := newFakeBalance()
	b2 := newFakeBalance()
	b2.AccountId = 99
	_, _ = service.Create(b1.AccountId)
	_, _ = service.Create(b2.AccountId)

	arr, err := service.List()
	assert.Equal(t, 2, len(arr))
	assert.Nil(t, err)
}

func Test_Delete(t *testing.T) {
	repo := newInmem()
	service := NewService(repo)
	b1 := newFakeBalance()
	b2 := newFakeBalance()
	b2.AccountId = 99
	i1, _ := service.Create(b1.AccountId)
	_, _ = service.Create(b2.AccountId)

	err := service.Delete(i1)
	assert.Nil(t, err)
	l, _ := service.List()
	assert.Equal(t, 1, len(l))
	g, err := service.Get(i1)
	assert.Nil(t, g)
	assert.NotNil(t, err)
	assert.Equal(t, domain.ErrNotFound, err)
}

func TestService_TransferFounds(t *testing.T) {
	repo := newInmem()
	service := NewService(repo)
	b1 := newFakeBalance()
	b2 := newFakeBalance()
	b2.AccountId = 99
	i1, _ := service.Create(b1.AccountId)
	i2, _ := service.Create(b2.AccountId)
	b1.ID = i1
	b2.ID = i2
	b1.Balance = 10
	b2.Balance = 10
	err := service.TransferFounds(b1, b2, 5)
	assert.Nil(t, err)
	up1, _ := service.Get(i1)
	up2, _ := service.Get(i2)
	assert.Equal(t, float32(5), up1.Balance)
	assert.Equal(t, float32(15), up2.Balance)

	err = service.TransferFounds(b1, b2, 3.5)
	assert.Nil(t, err)
	up1, _ = service.Get(i1)
	up2, _ = service.Get(i2)
	assert.Equal(t, float32(1.5), up1.Balance)
	assert.Equal(t, float32(18.5), up2.Balance)

	err = service.TransferFounds(b1, b2, 10)
	assert.NotNil(t, err)
	assert.Equal(t, domain.ErrInvalidBalance, err)
	assert.Equal(t, float32(1.5), up1.Balance)
	assert.Equal(t, float32(18.5), up2.Balance)

	err = service.TransferFounds(b1, b1, 10)
	assert.NotNil(t, err)
}
