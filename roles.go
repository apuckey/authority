package authority

// Role represents the database model of roles
type Role struct {
	ID   uint   `gorm:"AUTO_INCREMENT;primaryKey;column:id"`
	Name string `gorm:"column:name;unique_index:idx_role_name"`
}

// TableName sets the table name
func (r Role) TableName() string {
	return tablePrefix + "roles"
}
