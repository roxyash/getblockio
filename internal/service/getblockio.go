package service

import "github.com/roxyash/getblock/internal/types"

	type GetBlockIoService struct {
	}

	func NewGetBlockIoService() *GetBlockIoService {
		return &GetBlockIoService{}
	}

	func (s *GetBlockIoService) CheckBalance(apikey string) (types.ResponseCheckBalanceInfo, error) {
		return types.ResponseCheckBalanceInfo{}, nil 
	}