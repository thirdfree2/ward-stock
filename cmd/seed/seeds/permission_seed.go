package seeds

import (
	"ward-stock-backend/internal/domain"

	"gorm.io/gorm"
)

func SeedPermissions(db *gorm.DB) {
	perms := []domain.Permission{
		{Code: "user:create"},
		{Code: "user:read"},
		{Code: "user:update"},
		{Code: "user:delete"},
	}
	for _, p := range perms {
		db.FirstOrCreate(&p, domain.Permission{Code: p.Code})
	}
}
