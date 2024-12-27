package transaction

import "be-bwa-startup/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User `json:"user"` 
}