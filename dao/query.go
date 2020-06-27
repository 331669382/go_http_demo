package dao

func GetValByUsername(username string) (*User, error) {
	user := User{}
	err := db.
		Where("`username` = ?", username).
		First(&user).Error

	return &user, err
}
