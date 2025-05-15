package paymentservice

import (
	"context"
	paymentmodel "dev-coffee-api/modules/payments/model"
)

type GetPaymentByIDStorage interface {
	GetPaymentByID(ctx context.Context, id int) (*paymentmodel.Payment, error)
}

type GetPaymentByIDService struct {
	store GetPaymentByIDStorage
}

func NewGetPaymentByIDService(store GetPaymentByIDStorage) *GetPaymentByIDService {
	return &GetPaymentByIDService{store: store}
}

func (s *GetPaymentByIDService) GetPaymentByID(ctx context.Context, id int) (*paymentmodel.Payment, error) {
	return s.store.GetPaymentByID(ctx, id)
}
