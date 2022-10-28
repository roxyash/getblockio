package service

import "github.com/roxyash/getblock/internal/types"

type GetBlockIo interface {
	CheckBalance(apikey string) (types.ResponseCheckBalanceInfo, error)
}

type Service struct {
	GetBlockIo
}


func NewService() *Service {
	return &Service{
		GetBlockIo:  NewGetBlockIoService(),
	}
}
//go:generate mockgen -source=service.go -destination=mocks/mock.go
