package handler

import (
	"be-bwa-startup/helper"
	"be-bwa-startup/transaction"
	"be-bwa-startup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		respone := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignID(input)

	if err != nil {
		respone := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	response := helper.APIResponse("Transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))

	c.JSON(http.StatusOK, response)

}