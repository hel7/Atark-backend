package farmsage

type User struct {
	UserID   int    `json:"userID" db:"UserID"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" db:"role"`
}
