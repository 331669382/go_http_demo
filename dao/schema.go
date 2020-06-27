package dao

const (
	TableName_User = "users"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}
