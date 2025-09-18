package converter

import (
	"crispypay.com/challenge/api/contract"
	"crispypay.com/challenge/models"
	"github.com/google/uuid"
)

func ToPartnerRequest(request contract.TransferRequest) contract.PartnerRequest {
	return contract.PartnerRequest{
		ExternalID:  uuid.New().String(),
		Amount:      request.Amount,
		Currency:    request.Currency,
		Sender:      request.Sender,
		Receiver:    request.Receiver,
		CallbackURL: "http://localhost:8080/api/v1/webhooks/partner-callback",
	}

}

func ToTransferEntity(payload contract.TransferRequest, request contract.PartnerRequest, response contract.PartnerResponse) models.Transfer {
	return models.Transfer{
		ExternalID:           request.ExternalID,
		Amount:               request.Amount,
		Currency:             request.Currency,
		DestinationCountry:   payload.DestinationCountry,
		PartnerProvider:      payload.PartnerProvider,
		Purpose:              payload.Purpose,
		CallbackURL:          request.CallbackURL,
		PartnerTransactionID: response.PartnerTransactionID,
		Status:               response.Status,
		EstimatedCompletion:  response.EstimatedCompletion,
		Fee:                  response.Fee,
		Message:              response.Message,
	}
}
func ToUserEntity(userInfo contract.UserInfo, partnerTransactionID string, transactionType string) models.User {
	return models.User{
		PartnerTransactionID: partnerTransactionID,
		TransactionType:      transactionType,
		Name:                 userInfo.Name,
		Phone:                userInfo.Phone,
		Address:              userInfo.Address,
	}

}

func ToTransferHistoryEntity(transfer *models.Transfer) models.StatusHistory {
	return models.StatusHistory{
		TransferID: transfer.ID,
		Status:     transfer.Status,
	}
}

func ToSearchTransferResponse(transfer *models.Transfer, sender *models.User, receiver *models.User) contract.SearchTransferResponse {
	senderInfo := contract.UserInfo{
		Name:    sender.Name,
		Phone:   sender.Phone,
		Address: sender.Address,
	}

	receiverInfo := contract.UserInfo{
		Name:    receiver.Name,
		Phone:   receiver.Phone,
		Address: receiver.Address,
	}

	return contract.SearchTransferResponse{
		Amount:      transfer.Amount,
		Currency:    transfer.Currency,
		Sender:      senderInfo,
		Receiver:    receiverInfo,
		Status:      transfer.Status,
		Fee:         transfer.Fee,
		CompletedAt: transfer.CompletedAt,
	}
}

func ToRecentTransferResponses(transferJoined []contract.TransferJoined) []contract.SearchTransferResponse {
	var responses []contract.SearchTransferResponse = make([]contract.SearchTransferResponse, 0, len(transferJoined))

	for _, r := range transferJoined {
		response := contract.SearchTransferResponse{
			PartnerTransactionID: r.PartnerTransactionID,
			Status:               r.Status,
			Amount:               r.Amount,
			Currency:             r.Currency,
			CreatedAt:            r.CreatedAt,
			CompletedAt:          r.CompletedAt,
			Sender: contract.UserInfo{
				Name:    r.SenderName,
				Phone:   r.SenderPhone,
				Address: r.SenderAddress,
			},
			Receiver: contract.UserInfo{
				Name:    r.ReceiverName,
				Phone:   r.ReceiverPhone,
				Address: r.ReceiverAddress,
			},
		}
		responses = append(responses, response)
	}

	return responses
}
