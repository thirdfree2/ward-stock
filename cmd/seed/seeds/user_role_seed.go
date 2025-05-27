package seeds

import (
	"log"

	"ward-stock-backend/internal/domain"

	"gorm.io/gorm"
)

func SeedUserRoles(db *gorm.DB) {
	users := []struct {
		Email string
		Role  string
	}{
		{"admin@example.com", "admin"},
		{"developer@example.com", "developer"},
		{"member@example.com", "member"},
	}

	for _, entry := range users {
		var user domain.User
		var role domain.Role

		if err := db.Where("email = ?", entry.Email).First(&user).Error; err != nil {
			log.Printf("❌ User not found: %s", entry.Email)
			continue
		}

		if err := db.Where("name = ?", entry.Role).First(&role).Error; err != nil {
			log.Printf("❌ Role not found: %s", entry.Role)
			continue
		}

		if err := db.FirstOrCreate(&domain.UserRole{
			UserID: user.ID,
			RoleID: role.ID,
		}, domain.UserRole{
			UserID: user.ID,
			RoleID: role.ID,
		}).Error; err != nil {
			log.Printf("❌ Failed to assign role %s to user %s: %v", entry.Role, entry.Email, err)
		}
	}
}
