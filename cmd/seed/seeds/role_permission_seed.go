package seeds

import (
	"ward-stock-backend/internal/domain"

	"gorm.io/gorm"
)

func SeedRolePermissions(db *gorm.DB) {
	// === 1. ดึงทุก Permission ===
	var allPerms []domain.Permission
	db.Find(&allPerms)

	// === 2. ให้ admin + developer ได้เหมือนกัน ===
	roles := []string{"admin", "developer"}

	for _, roleName := range roles {
		var role domain.Role
		db.Where("name = ?", roleName).First(&role)

		for _, p := range allPerms {
			db.FirstOrCreate(&domain.RolePermission{
				RoleID:       role.ID,
				PermissionID: p.ID,
			}, domain.RolePermission{
				RoleID:       role.ID,
				PermissionID: p.ID,
			})
		}
	}

	// === 3. member ได้แค่ user:read ===
	var member domain.Role
	db.Where("name = ?", "member").First(&member)

	var readPerm domain.Permission
	db.Where("code = ?", "user:read").First(&readPerm)

	db.FirstOrCreate(&domain.RolePermission{
		RoleID:       member.ID,
		PermissionID: readPerm.ID,
	}, domain.RolePermission{
		RoleID:       member.ID,
		PermissionID: readPerm.ID,
	})
}
