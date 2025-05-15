package paymentmodel

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type PaymentMethod int

const (
	PaymentMethodCash PaymentMethod = iota
	PaymentMethodCard
	PaymentMethodBankTransfer
)

var allPaymentMethods = [3]string{"cash", "card", "bank_transfer"}

func (method *PaymentMethod) String() string {
	return allPaymentMethods[*method]
}

func parsePaymentMethod(s string) (PaymentMethod, error) {
	for i, v := range allPaymentMethods {
		if v == s {
			return PaymentMethod(i), nil
		}
	}

	return PaymentMethod(0), errors.New("invalid payment method")
}

func (method *PaymentMethod) Value() (driver.Value, error) {
	if method == nil {
		return nil, nil
	}

	return method.String(), nil
}

func (method *PaymentMethod) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}

	paymentMethod, err := parsePaymentMethod(string(bytes))
	if err != nil {
		return err
	}

	*method = paymentMethod
	return nil
}

func (method *PaymentMethod) MarshalJSON() ([]byte, error) {
	if method == nil {
		return nil, nil
	}

	return []byte(method.String()), nil
}

func (method *PaymentMethod) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	paymentMethod, err := parsePaymentMethod(str)
	if err != nil {
		return err
	}

	*method = paymentMethod
	return nil
}
