package services

import (
	"fmt"
	"net/http"
	"strings"

	"crispypay.com/challenge/api/contract"
	"crispypay.com/challenge/api/converter"
	"crispypay.com/challenge/infra"
	"crispypay.com/challenge/repositories"
	"crispypay.com/challenge/util"
)

type TransferService struct {
	transferRepository        repositories.TransferRepository
	userRepository            repositories.UserRepository
	transferHistoryRepository repositories.StatusHistoryRepository
	httpClient                *http.Client
}

func NewTransferService(
	transferRepository repositories.TransferRepository,
	userRepository repositories.UserRepository,
	transferHistoryRepository repositories.StatusHistoryRepository,
	client *http.Client,
) *TransferService {
	return &TransferService{
		transferRepository:        transferRepository,
		transferHistoryRepository: transferHistoryRepository,
		userRepository:            userRepository,
		httpClient:                client,
	}
}

// 파트너사로의 송금요청 및 해당 이력 기록
// 1. 파트너사로 송금 요청
// 2. 파트너사의 응답과 request 바탕으로 transfer 와 user 데이터 기록
func (s *TransferService) Transfer(request contract.TransferRequest) error {
	partnerRequest, partnerResponse, err := s.executePartnerTransfer(request)
	if err != nil {
		return err
	}

	transfer := converter.ToTransferEntity(request, partnerRequest, partnerResponse)
	sender := converter.ToUserEntity(partnerRequest.Sender, partnerResponse.PartnerTransactionID, "SENDER")
	receiver := converter.ToUserEntity(partnerRequest.Receiver, partnerResponse.PartnerTransactionID, "RECEIVER")

	if err := s.transferRepository.CreateTransfer(&transfer); err != nil {
		return fmt.Errorf("create transfer: %w", err)
	}
	if err := s.userRepository.CreateUser(&sender); err != nil {
		return fmt.Errorf("create sender: %w", err)
	}
	if err := s.userRepository.CreateUser(&receiver); err != nil {
		return fmt.Errorf("create receiver: %w", err)
	}

	transferHistory := converter.ToTransferHistoryEntity(&transfer)
	s.transferHistoryRepository.CreateStatusHistory(&transferHistory)

	return nil
}

// 파트너사 트랜잭션 ID 로 단건 조회
func (s *TransferService) SearchTransfer(partnerTxID string) contract.SearchTransferResponse {
	transfer := s.transferRepository.FindByPartnerTransactionID(partnerTxID)
	sender := s.userRepository.FindUser(partnerTxID, "SENDER")
	receiver := s.userRepository.FindUser(partnerTxID, "RECEIVER")

	return converter.ToSearchTransferResponse(transfer, sender, receiver)
}

// 최근 거래 목록 100개 조회
func (s *TransferService) SearchTransfers() []contract.SearchTransferResponse {
	transferJoinedList, _ := s.transferRepository.FindRecentWithUsers(100)

	return converter.ToRecentTransferResponses(transferJoinedList)
}

// PENDING 상태라면 거래 취소
func (s *TransferService) Cancel(partnerTxID string) error {
	transfer := s.transferRepository.FindByPartnerTransactionID(partnerTxID)
	upperStatus := strings.ToUpper(transfer.Status)

	if upperStatus == "PENDING" {
		transfer.Status = "CANCEL"
	} else {
		return nil
	}

	s.transferRepository.Save(transfer)

	transferHistory := converter.ToTransferHistoryEntity(transfer)
	s.transferHistoryRepository.CreateStatusHistory(&transferHistory)

	return nil
}

// 콜백을 받아 거래 완료 처리
func (s *TransferService) CompleteTransfer(callback contract.PartnerCallback) error {
	transfer := s.transferRepository.FindByPartnerTransactionID(callback.PartnerTransactionID)

	if transfer.Status == "PROCESSING" && callback.Status == "COMPLETED" {
		transfer.Status = "COMPLETED"
		transfer.CompletedAt = callback.CompletedAt
	} else {
		return nil
	}

	s.transferRepository.Save(transfer)

	transferHistory := converter.ToTransferHistoryEntity(transfer)
	s.transferHistoryRepository.CreateStatusHistory(&transferHistory)

	return nil
}

// 파트너사로의 송금요청
func (s *TransferService) executePartnerTransfer(request contract.TransferRequest) (contract.PartnerRequest, contract.PartnerResponse, error) {
	var partnerResponse contract.PartnerResponse
	partnerRequest := converter.ToPartnerRequest(request)

	url := makeURL(request.PartnerProvider)

	err := util.PostJSON(s.httpClient, url, partnerRequest, &partnerResponse)

	return partnerRequest, partnerResponse, err
}

// 파트너사 URL
func makeURL(partnerProvider string) string {
	upperProvider := strings.ToUpper(partnerProvider)

	port := ""

	switch upperProvider {
	case "FASTPAY":
		port = ":8001"
	case "SAFETRANSFER":
		port = ":8002"
	case "BUDGETSEND":
		port = ":8003"
	default:
		port = ":8001"
	}

	return "http://" + infra.GetHost() + port + "/api/transfer"
}
