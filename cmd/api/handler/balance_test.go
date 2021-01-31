package handler

import (
	"TransactionTest/internal/domain"
	"TransactionTest/internal/service/balance/mocks"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_listBalances(t *testing.T) {
	type mockBehavior func(r *mock_balance.MockUseCase, balances []*domain.Balance)

	createdAt := time.Now()

	tests := []struct {
		name                 string
		inputBody            string
		balances             []*domain.Balance
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: ``,
			balances: []*domain.Balance{
				{
					ID:        1,
					AccountId: 1,
					Balance:   10,
					CreatedAt: createdAt,
					UpdatedAt: createdAt,
				},
				{
					ID:        2,
					AccountId: 2,
					Balance:   5,
					CreatedAt: createdAt,
					UpdatedAt: createdAt,
				},
			},
			mockBehavior: func(r *mock_balance.MockUseCase, balances []*domain.Balance) {
				r.EXPECT().List().Return(balances, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: fmt.Sprintf(`[{"id":1,"account_id":1,"balance":10,"created_at":"%s","updated_at":"%s"},{"id":2,"account_id":2,"balance":5,"created_at":"%s","updated_at":"%s"}]`, createdAt.Format(time.RFC3339), createdAt.Format(time.RFC3339), createdAt.Format(time.RFC3339), createdAt.Format(time.RFC3339)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_balance.NewMockUseCase(c)

			tt.mockBehavior(s, tt.balances)

			handler := NewHandler(s)

			// Init Endpoint
			r := gin.New()
			r.GET("/balances/", func(c *gin.Context) {
			}, handler.listBalances)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/balances/", nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tt.expectedStatusCode)
			assert.Equal(t, w.Body.String(), tt.expectedResponseBody)
		})
	}
}

func TestHandler_makeTransaction(t *testing.T) {
	type mockBehavior func(r *mock_balance.MockUseCase, id1 domain.ID, id2 domain.ID, balances map[domain.ID]*domain.Balance, value float32)

	createdAt := time.Now()

	tests := []struct {
		name                 string
		inputBody            string
		balances             map[domain.ID]*domain.Balance
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"from_id": 1,"to_id": 2,"value": 5}`,
			balances: map[domain.ID]*domain.Balance{
				domain.ID(1): {
					ID:        1,
					AccountId: 1,
					Balance:   10,
					CreatedAt: createdAt,
					UpdatedAt: createdAt,
				},
				domain.ID(2): {
					ID:        2,
					AccountId: 2,
					Balance:   5,
					CreatedAt: createdAt,
					UpdatedAt: createdAt,
				},
			},
			mockBehavior: func(r *mock_balance.MockUseCase, id1 domain.ID, id2 domain.ID, balances map[domain.ID]*domain.Balance, value float32) {
				r.EXPECT().Get(id1).Return(balances[id1], nil)
				r.EXPECT().Get(id2).Return(balances[id2], nil)
				r.EXPECT().TransferFounds(balances[id1], balances[id2], value).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"status":"ok"}`,
		},
		{
			name:      "Ok",
			inputBody: `{"from_id": 1,"to_id": 2,"value": 5}`,
			balances: map[domain.ID]*domain.Balance{
				domain.ID(1): {
					ID:        1,
					AccountId: 1,
					Balance:   10,
					CreatedAt: createdAt,
					UpdatedAt: createdAt,
				},
			},
			mockBehavior: func(r *mock_balance.MockUseCase, id1 domain.ID, id2 domain.ID, balances map[domain.ID]*domain.Balance, value float32) {
				r.EXPECT().Get(id1).Return(balances[id1], nil)
				r.EXPECT().Get(id2).Return(nil, domain.ErrNotFound)
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"Not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_balance.NewMockUseCase(c)

			tt.mockBehavior(s, 1, 2, tt.balances, 5)

			handler := NewHandler(s)

			// Init Endpoint
			r := gin.New()
			r.POST("/balances/transaction", func(c *gin.Context) {
			}, handler.makeTransaction)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/balances/transaction", bytes.NewBufferString(tt.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tt.expectedStatusCode)
			assert.Equal(t, w.Body.String(), tt.expectedResponseBody)
		})
	}
}
