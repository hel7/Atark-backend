package farmsage

import (
	"errors"
	"regexp"
)

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

func (u *User) ValidateEmail() error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, u.Email)
	if !match {
		return errors.New("invalid email format")
	}
	return nil
}
func (u *User) ValidatePassword() error {
	if len(u.Password) < 8 {
		return errors.New("password should be at least 8 characters long")
	}

	letterRegex := regexp.MustCompile(`[a-zA-Z]`)
	if !letterRegex.MatchString(u.Password) {
		return errors.New("password should contain at least 1 letter")
	}

	digitRegex := regexp.MustCompile(`\d`)
	if !digitRegex.MatchString(u.Password) {
		return errors.New("password should contain at least 1 digit")
	}

	specialCharRegex := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	if !specialCharRegex.MatchString(u.Password) {
		return errors.New("password should contain at least 1 special character")
	}

	return nil
}
