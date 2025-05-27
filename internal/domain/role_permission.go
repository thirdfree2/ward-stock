package domain

type Permission struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique" json:"code"`
}

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique" json:"name"`
}

type RolePermission struct {
	ID           uint `gorm:"primaryKey"`
	RoleID       uint
	PermissionID uint

	Role       Role       `gorm:"foreignKey:RoleID"`
	Permission Permission `gorm:"foreignKey:PermissionID"`
}

type UserRole struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	RoleID uint

	User User `gorm:"foreignKey:UserID"`
	Role Role `gorm:"foreignKey:RoleID"`
}
