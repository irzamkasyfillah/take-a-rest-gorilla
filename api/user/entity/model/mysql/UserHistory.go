package mysql

// User represents a user.
type UserHistory struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Action string `json:"action"`
	Data   string `json:"data"`
}

// TableName returns the table name.
func UserHistoryTableName() string {
	return "users_history"
}
