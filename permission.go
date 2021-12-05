package authority

// Permission represents the database model of permissions
type Permission struct {
	ID   uint   `gorm:"AUTO_INCREMENT;primaryKey;column:id"`
	Name string `gorm:"column:name;unique_index:idx_permission_name"`
}

// TableName sets the table name
func (p Permission) TableName() string {
	return tablePrefix + "permissions"
}
