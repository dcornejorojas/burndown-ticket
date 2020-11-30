package models

import (
	"time"
)

/*Rule struct that describe if a product has certain business rule
@Format: Format of the store
@Store: Number of the store
@ProductID: Id of the product
@Type: Type of rule (decrease, high value, etc)
@Description: Legend that would be display with the description of the Rule
*/
type Rule struct {
	Format      int64 `json:"format"`
	Store       int64 `json:"store"`
	ProductID   string `json:"productId"`
	Type        int64 `json:"type"`
	Description string `json:"description"`
	Time        time.Time
}
