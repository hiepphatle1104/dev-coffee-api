package paymentstorage

import (
	"context"
	paymentmodel "dev-coffee-api/modules/payments/model"
)

func (s *sqlStorage) GetPaymentByID(ctx context.Context, id int) (*paymentmodel.Payment, error) {
	var data paymentmodel.Payment
	if err := s.db.Where("order_id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *sqlStorage) GetPaymentsList(ctx context.Context, paging *paymentmodel.Paging) (*[]paymentmodel.Payment, error) {
	var data []paymentmodel.Payment
	if err := s.db.Offset(paging.Offset()).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
