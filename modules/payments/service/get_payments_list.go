package paymentservice

import (
	paymentmodel "dev-coffee-api/modules/payments/model"
	"golang.org/x/net/context"
)

type GetPaymentsListStorage interface {
	GetPaymentsList(ctx context.Context, paging *paymentmodel.Paging) (*[]paymentmodel.Payment, error)
}

type GetPaymentsListService struct {
	store GetPaymentsListStorage
}

func NewGetPaymentsListService(store GetPaymentsListStorage) *GetPaymentsListService {
	return &GetPaymentsListService{store: store}
}

func (s *GetPaymentsListService) GetPaymentsList(ctx context.Context, paging *paymentmodel.Paging) (*[]paymentmodel.Payment, error) {
	return s.store.GetPaymentsList(ctx, paging)
}
