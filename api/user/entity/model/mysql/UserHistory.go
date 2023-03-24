package mysql

type UserHistoryData struct {
	Before map[string]interface{} `json:"before"`
	After  map[string]interface{} `json:"after"`
}
type UserHistory struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Action string `json:"action"`
	Data   string `json:"data"`
}

type UserHistoryRespond struct {
	ID     uint             `json:"id"`
	UserID uint             `json:"user_id"`
	Action string           `json:"action"`
	Data   *UserHistoryData `json:"data"`
}

// TableName returns the table name.
func UserHistoryTableName() string {
	return "users_history"
}
