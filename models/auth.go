package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	// todo::
	if username == "admin" && password == "admin" {
		return true
	}

	return false
}
