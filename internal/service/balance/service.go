package balance

import (
	"TransactionTest/internal/domain"
	"errors"
	"time"
)

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

//Service - Balance use case
type Service struct {
	repo Repository
}

//NewService return new Service
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

//Create a new domain.Balance
func (s *Service) Create(acId domain.ID) (domain.ID, error) {
	b, err := domain.NewBalance(acId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(b)
}

//Get a balance
func (s *Service) Get(id domain.ID) (*domain.Balance, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

//Update balance
func (s *Service) Update(b *domain.Balance) error {
	err := b.Validate()
	if err != nil {
		return err
	}
	b.UpdatedAt = time.Now()
	return s.repo.Update(b)
}

//Delete balance
func (s *Service) Delete(id domain.ID) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

//List returns all balances
func (s *Service) List() ([]*domain.Balance, error) {
	return s.repo.List()
}

//TransferFounds make balance transaction
func (s *Service) TransferFounds(t *domain.Balance, r *domain.Balance, val float32) error {
	if t.ID == r.ID {
		return errors.New("Can't make transaction to yourself")
	}
	err := t.ChangeBalance(-val)
	if err != nil {
		return err
	}
	err = r.ChangeBalance(val)
	if err != nil {
		return err
	}
	err = s.Update(t)
	if err != nil {
		return err
	}
	err = s.Update(r)
	if err != nil {
		return err
	}

	return nil
}
