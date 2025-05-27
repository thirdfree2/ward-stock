package seeds

import (
	"ward-stock-backend/internal/domain"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []domain.User{
		{
			Code:  "USR00001",
			Name:  "Admin User",
			Email: "admin@example.com",
		},
		{
			Code:  "USR00002",
			Name:  "Developer User",
			Email: "developer@example.com",
		},
		{
			Code:  "USR00003",
			Name:  "Member User",
			Email: "member@example.com",
		},
	}

	for _, u := range users {
		db.FirstOrCreate(&u, domain.User{Email: u.Email})
	}
}
