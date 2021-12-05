package authority

// UserRole represents the relationship between users and roles
type UserRole struct {
	ID     uint         `gorm:"AUTO_INCREMENT;primaryKey;column:id"`
	UserID UserIDColumn `gorm:"column:user_id;unique_index:idx_user_role"`
	RoleID uint         `gorm:"column:role_id;unique_index:idx_user_role"`
}

// TableName sets the table name
func (u UserRole) TableName() string {
	return tablePrefix + "user_roles"
}
