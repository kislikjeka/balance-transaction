package handler

import (
	"TransactionTest/cmd/api/presenter"
	"TransactionTest/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getBalanceById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	b, err := h.balanceService.Get(domain.ID(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if b == nil {
		newErrorResponse(c, http.StatusNotFound, "Balance not found")
		return
	}

	toJ := presenter.NewBalanceFromDomain(b)

	c.JSON(http.StatusOK, toJ)
}

func (h *Handler) updateBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input struct {
		ID        domain.ID `json:"id"`
		AccountID domain.ID `json:"account_id"`
		Balance   float32   `json: "balance"`
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	b := &domain.Balance{
		ID:        domain.ID(id),
		AccountId: input.AccountID,
		Balance:   input.Balance,
	}

	err = h.balanceService.Update(b)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.balanceService.Delete(domain.ID(id))
	if err !=nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) createBalance(c *gin.Context) {
	var input struct {
		AccountID domain.ID `json:"account_id"`
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.balanceService.Create(input.AccountID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) listBalances(c *gin.Context) {
	balances, err := h.balanceService.List()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var toJ []*presenter.Balance
	for _, b := range balances {
		toJ = append(toJ, presenter.NewBalanceFromDomain(b))
	}

	c.JSON(http.StatusOK, toJ)
}

func (h *Handler) makeTransaction(c *gin.Context) {
	var input struct {
		From domain.ID `json:"from_id"`
		To domain.ID `json:"to_id"`
		Value float32 `json:"value"`
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	t, err := h.balanceService.Get(input.From)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r, err := h.balanceService.Get(input.To)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.balanceService.TransferFounds(t,r,input.Value)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
