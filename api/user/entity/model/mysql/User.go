package mysql

// User represents a user.
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TableName returns the table name.
func UserTableName() string {
	return "users"
}
