package service

import (
	"ch13-seckill/oauth-service/model"
	"context"
	"errors"
)

var (
	ErrClientMessage = errors.New("invalid client")
)

// Service Define a service interface
type ClientDetailsService interface {
	GetClientDetailByClientId(ctx context.Context, clientId string, clientSecret string) (*model.ClientDetails, error)
}

type MysqlClientDetailsService struct {
}

func NewMysqlClientDetailsService() ClientDetailsService {
	return &MysqlClientDetailsService{}
}

func (service *MysqlClientDetailsService) GetClientDetailByClientId(ctx context.Context, clientId string, clientSecret string) (*model.ClientDetails, error) {

	clientDetailsModel := model.NewClientDetailsModel()
	if clientDetails, err := clientDetailsModel.GetClientDetailsByClientId(clientId); err == nil {
		if clientSecret == clientDetails.ClientSecret {
			return clientDetails, nil
		} else {
			return nil, ErrClientMessage
		}
	} else {
		return nil, err
	}

}
