package seeds

import (
	"ward-stock-backend/internal/domain"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	admin := domain.Role{Name: "admin"}
	db.FirstOrCreate(&admin, domain.Role{Name: "admin"})

	developer := domain.Role{Name: "developer"}
	db.FirstOrCreate(&developer, domain.Role{Name: "developer"})

	member := domain.Role{Name: "member"}
	db.FirstOrCreate(&member, domain.Role{Name: "member"})
}
