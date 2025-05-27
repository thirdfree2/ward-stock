// internal/domain/running_number.go
package domain

import (
	"ward-stock-backend/internal/domain/common"
)

type RunningNumber struct {
	ID     uint   `gorm:"primaryKey"`
	Key    string `gorm:"uniqueIndex"` // เช่น "PRODUCT", "INVOICE_20240527"
	LastNo int
	common.BaseModel
}

type RunningNumberRepository interface {
	NextNumber(key string, prefix string, digit int) (string, error)
}
