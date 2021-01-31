package postgres

import (
	"TransactionTest/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

const balanceTable = "balances"

//BalanceRepo - Postgres domain.Balance balance.Repository
type BalanceRepo struct {
	db *sqlx.DB
}

//NewBalanceRepo creates new BalanceRepo
func NewBalanceRepo(db *sqlx.DB) *BalanceRepo {
	return &BalanceRepo{
		db: db,
	}
}

//Create a new balance
func (r *BalanceRepo) Create(e *domain.Balance) (domain.ID, error) {
	var id domain.ID
	query := fmt.Sprintf("INSERT INTO %s (account_id, balance, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", balanceTable)

	row := r.db.QueryRow(query, e.AccountId, e.Balance, e.CreatedAt.Format(time.RFC3339), e.UpdatedAt.Format(time.RFC3339))
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

//Get balance by id
func (r *BalanceRepo) Get(id domain.ID) (*domain.Balance, error) {
	var balance domain.Balance
	query := fmt.Sprintf("SELECT id, account_id, balance, created_at, updated_at FROM %s WHERE id = $1", balanceTable)
	rows := r.db.QueryRowx(query, id)
	err := rows.Scan(
		&balance.ID,
		&balance.AccountId,
		&balance.Balance,
		&balance.CreatedAt,
		&balance.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &balance, nil
}

//List returns slice of balances
func (r *BalanceRepo) List() ([]*domain.Balance, error) {
	query := fmt.Sprintf("SELECT id, account_id, balance, created_at, updated_at FROM %s", balanceTable)
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	var balances []*domain.Balance
	for rows.Next() {
		var b domain.Balance
		err = rows.Scan(
			&b.ID,
			&b.AccountId,
			&b.Balance,
			&b.CreatedAt,
			&b.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		balances = append(balances, &b)
	}

	return balances, nil

}

//Update balance
func (r *BalanceRepo) Update(e *domain.Balance) error {
	e.UpdatedAt = time.Now()
	query := fmt.Sprintf("UPDATE %s SET balance = $1, updated_at = $2 WHERE id = $3", balanceTable)
	_, err := r.db.Exec(query, e.Balance, e.UpdatedAt.Format(time.RFC3339), e.ID)
	if err != nil {
		return err
	}

	return nil
}

//GetByAccountID get balance by account id
func (r *BalanceRepo) GetByAccountID(acid domain.ID) (*domain.Balance, error) {
	var balance domain.Balance
	query := fmt.Sprintf("SELECT id, account_id, balance, created_at, updated_at FROM %s WHERE account_id $1", balanceTable)
	rows, err := r.db.Queryx(query, acid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(
			&balance.ID,
			&balance.AccountId,
			&balance.Balance,
			&balance.CreatedAt,
			&balance.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	return &balance, nil
}

//Delete balance
func (r *BalanceRepo) Delete(id domain.ID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $4", balanceTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
