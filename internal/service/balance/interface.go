package balance

import "TransactionTest/internal/domain"

//Reader interface
type Reader interface {
	Get(id domain.ID) (*domain.Balance, error)
	GetByAccountID(id domain.ID) (*domain.Balance, error)
	List() ([]*domain.Balance, error)
}

//Writer book writer
type Writer interface {
	Create(e *domain.Balance) (domain.ID, error)
	Update(e *domain.Balance) error
	Delete(id domain.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	Get(id domain.ID) (*domain.Balance, error)
	Create(accountID domain.ID) (domain.ID, error)
	Update(b *domain.Balance) error
	Delete(id domain.ID) error
	List() ([]*domain.Balance, error)
	TransferFounds(t *domain.Balance, r *domain.Balance, val float32) error
}
