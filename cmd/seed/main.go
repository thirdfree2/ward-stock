package main

import (
	"log"

	"ward-stock-backend/cmd/seed/seeds"
	"ward-stock-backend/internal/domain"
	"ward-stock-backend/internal/infrastructure/postgres"
)

func main() {
	db := postgres.Connect()
	db.AutoMigrate(
		&domain.User{},
		&domain.Role{},
		&domain.Permission{},
		&domain.RolePermission{},
		&domain.UserRole{},
		&domain.RunningNumber{},
	)

	seeds.SeedPermissions(db)
	seeds.SeedRoles(db)
	seeds.SeedRolePermissions(db)
	seeds.SeedUsers(db)
	seeds.SetRunningNumberAfterUserSeed(db) // 👈 ตรงนี้
	seeds.SeedUserRoles(db)

	log.Println("✅ Seed completed.")
}
