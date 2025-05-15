package paymentservice

import (
	"context"
	ordermodel "dev-coffee-api/modules/orders/model"
	paymentmodel "dev-coffee-api/modules/payments/model"
)

type CreateNewPaymentStorage interface {
	CreateNewPayment(ctx context.Context, data *paymentmodel.PaymentCreation) error
	GetPaymentByID(ctx context.Context, id int) (*paymentmodel.Payment, error)
}

type CreateNewPaymentService struct {
	store          CreateNewPaymentStorage
	orderItemStore ordermodel.OrderItemStorage
}

func NewCreateNewPaymentService(store CreateNewPaymentStorage, orderItemStore ordermodel.OrderItemStorage) *CreateNewPaymentService {
	return &CreateNewPaymentService{store: store, orderItemStore: orderItemStore}
}

func (s *CreateNewPaymentService) CreateNewPayment(ctx context.Context, data *paymentmodel.PaymentCreation) error {
	_, err := s.store.GetPaymentByID(ctx, data.OrderID)
	if err != nil {
		return err
	}

	orderItems, err := s.orderItemStore.GetOrderItemsByID(ctx, data.OrderID)
	if err != nil {
		return err
	}

	// Summarize total amount
	var amount float64 = 0
	for _, item := range *orderItems {
		amount += item.UnitPrice * float64(item.Quantity)
	}

	data.Amount = amount

	return s.store.CreateNewPayment(ctx, data)
}
