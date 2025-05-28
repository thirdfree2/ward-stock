package model

type UserRoleModel struct {
	ID     uint
	UserID uint
	RoleID uint
	Role   RoleModel  `gorm:"foreignKey:RoleID"`
	User   *UserModel `gorm:"foreignKey:UserID"`
}

func (UserRoleModel) TableName() string {
	return "user_roles"
}
