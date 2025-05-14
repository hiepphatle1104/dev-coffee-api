package itemmodel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemType int

const (
	ItemTypeCoffee ItemType = iota
	ItemTypeTea
	ItemTypeJuice
	ItemTypeFood
)

var allItemTypes = [4]string{"coffee", "tea", "juice", "food"}

// This function use for turn type to string to insert into db
func (t *ItemType) String() string {
	return allItemTypes[*t]
}

func parseItemType(s string) (ItemType, error) {
	for i, v := range allItemTypes {
		if v == s {
			return ItemType(i), nil
		}
	}

	return ItemType(0), errors.New("invalid item type")
}

// Scan Data from db to client
func (t *ItemType) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("failed to scan data from sql: %s", value))
	}

	itemType, err := parseItemType(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintf("failed to scan data from sql: %s", value))
	}

	*t = itemType

	return nil
}

// Value Insert data to database
func (t *ItemType) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}

	return t.String(), nil
}

func (t *ItemType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *ItemType) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	itemType, err := parseItemType(str)
	if err != nil {
		return err
	}

	*t = itemType

	return nil
}
