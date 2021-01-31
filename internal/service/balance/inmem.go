package balance

import (
	"TransactionTest/internal/domain"
	"math/rand"
)

type inmem struct {
	m map[domain.ID]*domain.Balance
}

func newInmem() * inmem{
	m := make(map[domain.ID]*domain.Balance)
	return &inmem{
		m: m,
	}
}

func (r inmem) Get(id domain.ID) (*domain.Balance, error) {
	if r.m[id] == nil {
		return nil, domain.ErrNotFound
	}
	return r.m[id], nil
}


func (r inmem) GetByAccountID(id domain.ID) (*domain.Balance, error) {
	for _, bal := range r.m {
		if bal.AccountId == id {
			return bal, nil
		}
	}
	return nil, domain.ErrNotFound
}

func (r inmem) List() ([]*domain.Balance, error) {
	var d []*domain.Balance
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

func (r inmem) Create(e *domain.Balance) (domain.ID, error) {
	id := domain.ID(rand.Uint32())
	b, err := domain.NewBalance(e.AccountId)
	b.ID = id
	if err != nil {
		return 0, nil
	}
	r.m[id] = b
	return id, nil
}

func (r inmem) Update(e *domain.Balance) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

func (r inmem) Delete(id domain.ID) error {
	if r.m[id] == nil {
		return domain.ErrNotFound
	}
	delete(r.m, id)
	return nil
}

