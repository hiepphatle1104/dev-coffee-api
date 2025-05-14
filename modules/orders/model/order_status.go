package ordermodel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type OrderStatus int

const (
	OrderStatusPending OrderStatus = iota
	OrderStatusCancelled
	OrderStatusCompleted
)

var allOrderStatus = [3]string{"pending", "cancelled", "completed"}

func (s *OrderStatus) String() string {
	return allOrderStatus[*s]
}

func (s *OrderStatus) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}

	return s.String(), nil
}

func (s *OrderStatus) MarshalJSON() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", s.String())), nil
}

func (s *OrderStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	orderStatus, err := parseOrderStatus(str)
	if err != nil {
		return err
	}

	*s = orderStatus
	return nil
}

func parseOrderStatus(s string) (OrderStatus, error) {
	for i, v := range allOrderStatus {
		if v == s {
			return OrderStatus(i), nil
		}
	}

	return OrderStatus(0), errors.New("invalid order status")
}

func (s *OrderStatus) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("failed to scan data from sql: %s", value))
	}

	orderStatus, err := parseOrderStatus(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintf("failed to scan data from sql: %s", value))
	}

	*s = orderStatus

	return nil
}
