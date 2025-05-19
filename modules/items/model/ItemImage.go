package itemmodel

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ItemImage struct {
	URL    string `json:"url"`
	Alt    string `json:"alt"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (i *ItemImage) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (i *ItemImage) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan ItemImage: %v", value)
	}
	return json.Unmarshal(bytes, i)
}
