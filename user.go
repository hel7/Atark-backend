package farmsage

import "errors"

type User struct {
	UserID   int    `json:"userID" db:"UserID"`
	Username string `json:"username" binding:"required" db:"Username"`
	Email    string `json:"email" binding:"required" db:"Email"`
	Password string `json:"password" binding:"required" db:"Password"`
	Role     string `json:"user_role" db:"Role"`
}

type UpdateUserInput struct {
	UserID   *int    `json:"UserID"`
	Username *string `json:"Username"`
	Email    *string `json:"Email"`
	Password *string `json:"Password"`
	Role     *string `json:"Role"`
}

func (i UpdateUserInput) Validate() error {
	if i.Username == nil && i.Email == nil && i.Password == nil && i.Role == nil {
		return errors.New("Update structure has no value")
	}
	return nil
}
