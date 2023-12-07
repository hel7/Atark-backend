package farmsage

type User struct {
	UserID   int    `json:"userID" db:"UserID"`
	Username string `json:"username" binding:"required" db:"Username"`
	Email    string `json:"email" binding:"required" db:"Email"`
	Password string `json:"password" binding:"required" db:"Password"`
	Role     string `json:"user_role" db:"Role"`
}
