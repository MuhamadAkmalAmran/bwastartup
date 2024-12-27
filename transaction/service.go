package transaction

import (
	"be-bwa-startup/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
}

func NewService(repository Repository, campainRepository campaign.Repository) *service {
	return &service{repository, campainRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindById(input.ID)

	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transaction, err := s.repository.GetByCampaignID(input.ID)

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}